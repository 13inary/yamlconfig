package yamlconfig

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	defaultConfigFile = "config.yaml"

	defaultTag = "default"
)

// yamlLoader 配置加载器
type yamlLoader struct {
	configFile        string
	exampleConfigFile string
}

// LoadYamlConfig 加载YAML配置文件
// configFile: 配置文件路径，若不设置tag，那么配置文件中字段需要全部小写
// exampleConfigFile: 示例配置文件路径
// config: 必须是指向配置结构体的指针。yaml标签指定配置文件中字段名（更新字段名时记得同步修改），default标签指定默认值
func LoadYamlConfig(configFile, exampleConfigFile string, config any) error {
	loader := NewYamlLoader(configFile, exampleConfigFile)
	return loader.Load(config)
}

// NewYamlLoader 创建新的配置加载器
// configFile: 配置文件路径，若不设置tag，那么配置文件中字段需要全部小写
// exampleConfigFile: 示例配置文件路径
func NewYamlLoader(configFile, exampleConfigFile string) *yamlLoader {
	if configFile == "" {
		configFile = defaultConfigFile
	}
	return &yamlLoader{
		configFile:        configFile,
		exampleConfigFile: exampleConfigFile,
	}
}

// Load 加载配置文件到指定的配置结构体
// config: 必须是指向配置结构体的指针。yaml标签指定配置文件中字段名（更新字段名时记得同步修改），default标签指定默认值
func (l *yamlLoader) Load(config any) error {
	if config == nil {
		return fmt.Errorf("config is nil")
	}

	rv := reflect.ValueOf(config)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("config must be a non-nil pointer to struct")
	}

	configType := rv.Type().Elem()

	// 如果指定了示例配置文件路径，且示例配置文件不存在，则创建示例配置文件
	if l.exampleConfigFile != "" {
		if err := l.createExampleConfig(configType, l.exampleConfigFile); err != nil {
			return err
		}
	}

	// 读取并解析配置文件（如果存在）
	if err := l.loadConfigFile(config); err != nil {
		return err
	}

	// 设置默认值（仅在字段为零值时设置）
	if err := l.fillStructFromTag(config, true); err != nil {
		return fmt.Errorf("failed to set defaults: %w", err)
	}

	return nil
}

// fileExists 检查文件是否存在
func (l *yamlLoader) fileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// // ensureExampleConfig 确保示例配置文件存在
// func (l *yamlLoader) ensureExampleConfig(configType reflect.Type) error {
// 	exists, err := l.fileExists(l.exampleConfigFile)
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		// 示例文件不存在，创建示例配置文件
// 		return l.createExampleConfig(configType, l.exampleConfigFile)
// 	}
// 	return nil
// }

// loadConfigFile 读取并解析配置文件（如果存在）
func (l *yamlLoader) loadConfigFile(config any) error {
	exists, err := l.fileExists(l.configFile)
	if err != nil {
		return err
	}
	if !exists {
		// 配置文件不存在，跳过读取
		return nil
	}

	// 读取配置文件
	data, err := os.ReadFile(l.configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// 解析配置文件到结构体
	if err := yaml.Unmarshal(data, config); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	return nil
}

// createExampleConfig 创建示例默认配置文件。若已经存在旧的，则直接覆盖旧的
func (l *yamlLoader) createExampleConfig(configType reflect.Type, filePath string) error {
	// 创建配置实例
	config := reflect.New(configType).Interface()

	// 填充所有默认值（用于生成示例配置）
	if err := l.fillStructFromTag(config, false); err != nil {
		return fmt.Errorf("failed to fill defaults from tags: %w", err)
	}

	// 将配置序列化为YAML
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// 添加注释头
	// 预分配容量以提高性能
	headerLen := len("# 示例配置文件\n# 复制此文件为 ") + len(l.configFile) + len(" 并根据需要修改配置\n\n")
	var sb strings.Builder
	sb.Grow(headerLen + len(data))
	sb.WriteString("# 示例配置文件\n")
	sb.WriteString("# 复制此文件为 ")
	sb.WriteString(l.configFile)
	sb.WriteString(" 并根据需要修改配置\n\n")
	sb.Write(data)

	// 写入文件（没有做sync操作）
	// 注意：0644 是 Unix 权限位，在 Windows 上会被忽略（这是正常行为）
	return os.WriteFile(filePath, []byte(sb.String()), 0644)
}

// fillStructFromTag 处理结构体，设置默认值
// onlyIfZero: true表示仅在字段为零值时设置默认值，false表示总是设置（用于生成示例配置）
func (l *yamlLoader) fillStructFromTag(v any, onlyIfZero bool) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("fillStructFromTag: v must be a non-nil pointer")
	}

	rv = rv.Elem()
	rt := rv.Type()

	// 处理结构体字段
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		fieldType := rt.Field(i)

		if err := l.processField(field, fieldType, onlyIfZero); err != nil {
			return err
		}
	}

	return nil
}

// isStructPointerType 检查类型是否为结构体指针类型
func (l *yamlLoader) isStructPointerType(typ reflect.Type) bool {
	return typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct
}

// processField 处理单个字段
func (l *yamlLoader) processField(field reflect.Value, fieldType reflect.StructField, onlyIfZero bool) error {
	switch field.Kind() {
	case reflect.Slice: // 切片类型字段
		return l.processSliceField(field, fieldType, onlyIfZero)
	case reflect.Ptr: // 指针类型字段
		// 只支持结构体指针类型
		if !l.isStructPointerType(fieldType.Type) {
			return fmt.Errorf("field %s: pointer type must point to struct, got %s", fieldType.Name, fieldType.Type.Elem().Kind())
		}
		// 如果指针为nil，创建新的结构体实例（始终保证结构体指针不会是nil）
		if field.IsNil() {
			field.Set(reflect.New(fieldType.Type.Elem()))
		}
		// 递归处理结构体指针指向的结构体
		return l.fillStructFromTag(field.Interface(), onlyIfZero)
	case reflect.Struct: // 结构体类型字段，递归处理
		if field.CanAddr() {
			return l.fillStructFromTag(field.Addr().Interface(), onlyIfZero)
		}
		// 不可寻址的结构体字段（理论上不应该出现，但为了健壮性保留）
		return nil
	default: // 基本类型字段
		return l.processBasicField(field, fieldType, onlyIfZero)
	}
}

// processBasicField 处理基本类型字段
func (l *yamlLoader) processBasicField(field reflect.Value, fieldType reflect.StructField, onlyIfZero bool) error {
	// bool类型不支持default tag（避免零值歧义）
	if field.Kind() == reflect.Bool {
		return nil
	}

	// 有值时，跳过设置默认值
	if onlyIfZero && !field.IsZero() {
		return nil
	}

	defaultTag := fieldType.Tag.Get(defaultTag)
	if defaultTag == "" {
		// 若没有default tag的字段，保持原本值
		return nil
	}

	// 如果有default tag，解析并设置值
	if err := l.setFieldValue(field, defaultTag); err != nil {
		return fmt.Errorf("processBasicField: field %s: %w", fieldType.Name, err)
	}

	return nil
}

// processSliceField 处理切片类型字段
func (l *yamlLoader) processSliceField(field reflect.Value, fieldType reflect.StructField, onlyIfZero bool) error {
	// 如果切片为nil，初始化为空切片（始终保证切片不会是nil）
	if field.IsNil() {
		field.Set(reflect.MakeSlice(fieldType.Type, 0, 0))
	}

	// 无论onlyIfZero何值，都要处理切片的元素，因此切片的元素是结构体时，结构体的字段要处理默认值
	elemType := fieldType.Type.Elem()
	for i := 0; i < field.Len(); i++ {
		if err := l.processSliceElement(field.Index(i), elemType, onlyIfZero); err != nil {
			return fmt.Errorf("processSliceField: index %d: %w", i, err)
		}
	}
	return nil
}

// processSliceElement 处理切片中的单个元素
func (l *yamlLoader) processSliceElement(elem reflect.Value, elemType reflect.Type, onlyIfZero bool) error {
	switch elemType.Kind() {
	case reflect.Ptr:
		// 只支持结构体指针类型
		if !l.isStructPointerType(elemType) {
			return fmt.Errorf("slice element pointer type must point to struct, got %s", elemType.Elem().Kind())
		}
		// nil指针直接跳过（根据设计，切片中的nil指针元素不处理）
		if elem.IsNil() {
			return nil
		}
		return l.fillStructFromTag(elem.Interface(), onlyIfZero)

	case reflect.Struct:
		// 结构体类型元素，递归处理
		if elem.CanAddr() {
			return l.fillStructFromTag(elem.Addr().Interface(), onlyIfZero)
		}
		// 不可寻址的结构体元素（理论上不应该出现，但为了健壮性保留）
		return nil

	default:
		// 基本类型元素，YAML已解析，无需处理
		return nil
	}
}

// setFieldValue 根据字段类型设置值
func (l *yamlLoader) setFieldValue(field reflect.Value, value string) error {
	if !field.CanSet() {
		return fmt.Errorf("field cannot be set")
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse int value %q: %w", value, err)
		}
		field.SetInt(intVal)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintVal, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse uint value %q: %w", value, err)
		}
		field.SetUint(uintVal)

	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("failed to parse float value %q: %w", value, err)
		}
		field.SetFloat(floatVal)

	case reflect.Bool:
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("failed to parse bool value %q: %w", value, err)
		}
		field.SetBool(boolVal)

	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}

	return nil
}
