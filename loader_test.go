package yamlconfig

import (
	"os"
	"path/filepath"
	"testing"
)

// TestBasicTypesParsing 测试场景1：各种类型的字段是否解析正常
func TestBasicTypesParsing(t *testing.T) {
	// 定义测试配置结构体，包含各种基本数据类型
	type Config struct {
		StringField  string  `yaml:"string_field"`
		IntField     int     `yaml:"int_field"`
		Int8Field    int8    `yaml:"int8_field"`
		Int16Field   int16   `yaml:"int16_field"`
		Int32Field   int32   `yaml:"int32_field"`
		Int64Field   int64   `yaml:"int64_field"`
		UintField    uint    `yaml:"uint_field"`
		Uint8Field   uint8   `yaml:"uint8_field"`
		Uint16Field  uint16  `yaml:"uint16_field"`
		Uint32Field  uint32  `yaml:"uint32_field"`
		Uint64Field  uint64  `yaml:"uint64_field"`
		Float32Field float32 `yaml:"float32_field"`
		Float64Field float64 `yaml:"float64_field"`
		BoolField    bool    `yaml:"bool_field"`
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建测试用的 YAML 配置文件
	yamlContent := `string_field: "test string"
int_field: 42
int8_field: 8
int16_field: 16
int32_field: 32
int64_field: 64
uint_field: 100
uint8_field: 200
uint16_field: 300
uint32_field: 400
uint64_field: 500
float32_field: 3.14
float64_field: 2.71828
bool_field: true
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证各个字段是否正确解析
	if config.StringField != "test string" {
		t.Errorf("StringField 解析错误: 期望 %q, 实际 %q", "test string", config.StringField)
	}

	if config.IntField != 42 {
		t.Errorf("IntField 解析错误: 期望 %d, 实际 %d", 42, config.IntField)
	}

	if config.Int8Field != 8 {
		t.Errorf("Int8Field 解析错误: 期望 %d, 实际 %d", 8, config.Int8Field)
	}

	if config.Int16Field != 16 {
		t.Errorf("Int16Field 解析错误: 期望 %d, 实际 %d", 16, config.Int16Field)
	}

	if config.Int32Field != 32 {
		t.Errorf("Int32Field 解析错误: 期望 %d, 实际 %d", 32, config.Int32Field)
	}

	if config.Int64Field != 64 {
		t.Errorf("Int64Field 解析错误: 期望 %d, 实际 %d", 64, config.Int64Field)
	}

	if config.UintField != 100 {
		t.Errorf("UintField 解析错误: 期望 %d, 实际 %d", 100, config.UintField)
	}

	if config.Uint8Field != 200 {
		t.Errorf("Uint8Field 解析错误: 期望 %d, 实际 %d", 200, config.Uint8Field)
	}

	if config.Uint16Field != 300 {
		t.Errorf("Uint16Field 解析错误: 期望 %d, 实际 %d", 300, config.Uint16Field)
	}

	if config.Uint32Field != 400 {
		t.Errorf("Uint32Field 解析错误: 期望 %d, 实际 %d", 400, config.Uint32Field)
	}

	if config.Uint64Field != 500 {
		t.Errorf("Uint64Field 解析错误: 期望 %d, 实际 %d", 500, config.Uint64Field)
	}

	if config.Float32Field != 3.14 {
		t.Errorf("Float32Field 解析错误: 期望 %f, 实际 %f", 3.14, config.Float32Field)
	}

	if config.Float64Field != 2.71828 {
		t.Errorf("Float64Field 解析错误: 期望 %f, 实际 %f", 2.71828, config.Float64Field)
	}

	if config.BoolField != true {
		t.Errorf("BoolField 解析错误: 期望 %v, 实际 %v", true, config.BoolField)
	}

	t.Log("✓ 所有基本类型字段解析测试通过")
}

// TestNestedStructParsing 测试场景2：结构体的字段是结构体的情况下各种类型的字段是否解析正常
func TestNestedStructParsing(t *testing.T) {
	// 定义嵌套的结构体
	type DatabaseConfig struct {
		Host     string  `yaml:"host"`
		Port     int     `yaml:"port"`
		Username string  `yaml:"username"`
		Password string  `yaml:"password"`
		Timeout  float64 `yaml:"timeout"`
		Enabled  bool    `yaml:"enabled"`
	}

	type ServerConfig struct {
		Address string `yaml:"address"`
		Port    int    `yaml:"port"`
	}

	type Config struct {
		AppName string        `yaml:"app_name"`
		Version int           `yaml:"version"`
		Database DatabaseConfig `yaml:"database"`
		Server   ServerConfig `yaml:"server"`
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建测试用的 YAML 配置文件，包含嵌套结构体
	yamlContent := `app_name: "myapp"
version: 1
database:
  host: "localhost"
  port: 5432
  username: "admin"
  password: "secret123"
  timeout: 30.5
  enabled: true
server:
  address: "0.0.0.0"
  port: 8080
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证顶层字段
	if config.AppName != "myapp" {
		t.Errorf("AppName 解析错误: 期望 %q, 实际 %q", "myapp", config.AppName)
	}

	if config.Version != 1 {
		t.Errorf("Version 解析错误: 期望 %d, 实际 %d", 1, config.Version)
	}

	// 验证嵌套的 DatabaseConfig 结构体字段
	if config.Database.Host != "localhost" {
		t.Errorf("Database.Host 解析错误: 期望 %q, 实际 %q", "localhost", config.Database.Host)
	}

	if config.Database.Port != 5432 {
		t.Errorf("Database.Port 解析错误: 期望 %d, 实际 %d", 5432, config.Database.Port)
	}

	if config.Database.Username != "admin" {
		t.Errorf("Database.Username 解析错误: 期望 %q, 实际 %q", "admin", config.Database.Username)
	}

	if config.Database.Password != "secret123" {
		t.Errorf("Database.Password 解析错误: 期望 %q, 实际 %q", "secret123", config.Database.Password)
	}

	if config.Database.Timeout != 30.5 {
		t.Errorf("Database.Timeout 解析错误: 期望 %f, 实际 %f", 30.5, config.Database.Timeout)
	}

	if config.Database.Enabled != true {
		t.Errorf("Database.Enabled 解析错误: 期望 %v, 实际 %v", true, config.Database.Enabled)
	}

	// 验证嵌套的 ServerConfig 结构体字段
	if config.Server.Address != "0.0.0.0" {
		t.Errorf("Server.Address 解析错误: 期望 %q, 实际 %q", "0.0.0.0", config.Server.Address)
	}

	if config.Server.Port != 8080 {
		t.Errorf("Server.Port 解析错误: 期望 %d, 实际 %d", 8080, config.Server.Port)
	}

	t.Log("✓ 嵌套结构体字段解析测试通过")
}

// TestSliceFieldParsing 测试场景3：结构体的字段是切片的情况下各种类型的字段是否解析正常
func TestSliceFieldParsing(t *testing.T) {
	// 定义包含各种切片类型的配置结构体
	type Config struct {
		StringSlice  []string  `yaml:"string_slice"`
		IntSlice     []int     `yaml:"int_slice"`
		FloatSlice   []float64 `yaml:"float_slice"`
		BoolSlice    []bool    `yaml:"bool_slice"`
		UintSlice    []uint    `yaml:"uint_slice"`
		Int32Slice   []int32   `yaml:"int32_slice"`
		Float32Slice []float32 `yaml:"float32_slice"`
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建测试用的 YAML 配置文件，包含各种切片
	yamlContent := `string_slice:
  - "item1"
  - "item2"
  - "item3"
int_slice:
  - 1
  - 2
  - 3
  - 4
float_slice:
  - 1.1
  - 2.2
  - 3.3
bool_slice:
  - true
  - false
  - true
uint_slice:
  - 10
  - 20
  - 30
int32_slice:
  - 100
  - 200
float32_slice:
  - 1.5
  - 2.5
  - 3.5
  - 4.5
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证 StringSlice
	expectedStringSlice := []string{"item1", "item2", "item3"}
	if len(config.StringSlice) != len(expectedStringSlice) {
		t.Errorf("StringSlice 长度错误: 期望 %d, 实际 %d", len(expectedStringSlice), len(config.StringSlice))
	} else {
		for i, v := range expectedStringSlice {
			if config.StringSlice[i] != v {
				t.Errorf("StringSlice[%d] 解析错误: 期望 %q, 实际 %q", i, v, config.StringSlice[i])
			}
		}
	}

	// 验证 IntSlice
	expectedIntSlice := []int{1, 2, 3, 4}
	if len(config.IntSlice) != len(expectedIntSlice) {
		t.Errorf("IntSlice 长度错误: 期望 %d, 实际 %d", len(expectedIntSlice), len(config.IntSlice))
	} else {
		for i, v := range expectedIntSlice {
			if config.IntSlice[i] != v {
				t.Errorf("IntSlice[%d] 解析错误: 期望 %d, 实际 %d", i, v, config.IntSlice[i])
			}
		}
	}

	// 验证 FloatSlice
	expectedFloatSlice := []float64{1.1, 2.2, 3.3}
	if len(config.FloatSlice) != len(expectedFloatSlice) {
		t.Errorf("FloatSlice 长度错误: 期望 %d, 实际 %d", len(expectedFloatSlice), len(config.FloatSlice))
	} else {
		for i, v := range expectedFloatSlice {
			if config.FloatSlice[i] != v {
				t.Errorf("FloatSlice[%d] 解析错误: 期望 %f, 实际 %f", i, v, config.FloatSlice[i])
			}
		}
	}

	// 验证 BoolSlice
	expectedBoolSlice := []bool{true, false, true}
	if len(config.BoolSlice) != len(expectedBoolSlice) {
		t.Errorf("BoolSlice 长度错误: 期望 %d, 实际 %d", len(expectedBoolSlice), len(config.BoolSlice))
	} else {
		for i, v := range expectedBoolSlice {
			if config.BoolSlice[i] != v {
				t.Errorf("BoolSlice[%d] 解析错误: 期望 %v, 实际 %v", i, v, config.BoolSlice[i])
			}
		}
	}

	// 验证 UintSlice
	expectedUintSlice := []uint{10, 20, 30}
	if len(config.UintSlice) != len(expectedUintSlice) {
		t.Errorf("UintSlice 长度错误: 期望 %d, 实际 %d", len(expectedUintSlice), len(config.UintSlice))
	} else {
		for i, v := range expectedUintSlice {
			if config.UintSlice[i] != v {
				t.Errorf("UintSlice[%d] 解析错误: 期望 %d, 实际 %d", i, v, config.UintSlice[i])
			}
		}
	}

	// 验证 Int32Slice
	expectedInt32Slice := []int32{100, 200}
	if len(config.Int32Slice) != len(expectedInt32Slice) {
		t.Errorf("Int32Slice 长度错误: 期望 %d, 实际 %d", len(expectedInt32Slice), len(config.Int32Slice))
	} else {
		for i, v := range expectedInt32Slice {
			if config.Int32Slice[i] != v {
				t.Errorf("Int32Slice[%d] 解析错误: 期望 %d, 实际 %d", i, v, config.Int32Slice[i])
			}
		}
	}

	// 验证 Float32Slice
	expectedFloat32Slice := []float32{1.5, 2.5, 3.5, 4.5}
	if len(config.Float32Slice) != len(expectedFloat32Slice) {
		t.Errorf("Float32Slice 长度错误: 期望 %d, 实际 %d", len(expectedFloat32Slice), len(config.Float32Slice))
	} else {
		for i, v := range expectedFloat32Slice {
			if config.Float32Slice[i] != v {
				t.Errorf("Float32Slice[%d] 解析错误: 期望 %f, 实际 %f", i, v, config.Float32Slice[i])
			}
		}
	}

	// 验证切片不是 nil（根据 README 要求：对于切片，代码会始终保证不是nil切片）
	if config.StringSlice == nil {
		t.Error("StringSlice 不应该是 nil")
	}
	if config.IntSlice == nil {
		t.Error("IntSlice 不应该是 nil")
	}
	if config.FloatSlice == nil {
		t.Error("FloatSlice 不应该是 nil")
	}
	if config.BoolSlice == nil {
		t.Error("BoolSlice 不应该是 nil")
	}
	if config.UintSlice == nil {
		t.Error("UintSlice 不应该是 nil")
	}
	if config.Int32Slice == nil {
		t.Error("Int32Slice 不应该是 nil")
	}
	if config.Float32Slice == nil {
		t.Error("Float32Slice 不应该是 nil")
	}

	t.Log("✓ 切片字段解析测试通过")
}

// TestSliceOfStructParsing 测试场景4：结构体的字段是切片且切片的元素是结构体的情况下各种类型的字段是否解析正常
func TestSliceOfStructParsing(t *testing.T) {
	// 定义切片元素的结构体
	type User struct {
		Name  string `yaml:"name"`
		Age   int    `yaml:"age"`
		Email string `yaml:"email"`
		Score float64 `yaml:"score"`
		Active bool   `yaml:"active"`
	}

	type Config struct {
		Users []User `yaml:"users"`
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建测试用的 YAML 配置文件，包含结构体切片
	yamlContent := `users:
  - name: "Alice"
    age: 25
    email: "alice@example.com"
    score: 95.5
    active: true
  - name: "Bob"
    age: 30
    email: "bob@example.com"
    score: 87.0
    active: false
  - name: "Charlie"
    age: 28
    email: "charlie@example.com"
    score: 92.3
    active: true
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证切片长度
	expectedUsers := []User{
		{Name: "Alice", Age: 25, Email: "alice@example.com", Score: 95.5, Active: true},
		{Name: "Bob", Age: 30, Email: "bob@example.com", Score: 87.0, Active: false},
		{Name: "Charlie", Age: 28, Email: "charlie@example.com", Score: 92.3, Active: true},
	}

	if len(config.Users) != len(expectedUsers) {
		t.Fatalf("Users 切片长度错误: 期望 %d, 实际 %d", len(expectedUsers), len(config.Users))
	}

	// 验证每个结构体元素的字段
	for i, expectedUser := range expectedUsers {
		actualUser := config.Users[i]
		if actualUser.Name != expectedUser.Name {
			t.Errorf("Users[%d].Name 解析错误: 期望 %q, 实际 %q", i, expectedUser.Name, actualUser.Name)
		}
		if actualUser.Age != expectedUser.Age {
			t.Errorf("Users[%d].Age 解析错误: 期望 %d, 实际 %d", i, expectedUser.Age, actualUser.Age)
		}
		if actualUser.Email != expectedUser.Email {
			t.Errorf("Users[%d].Email 解析错误: 期望 %q, 实际 %q", i, expectedUser.Email, actualUser.Email)
		}
		if actualUser.Score != expectedUser.Score {
			t.Errorf("Users[%d].Score 解析错误: 期望 %f, 实际 %f", i, expectedUser.Score, actualUser.Score)
		}
		if actualUser.Active != expectedUser.Active {
			t.Errorf("Users[%d].Active 解析错误: 期望 %v, 实际 %v", i, expectedUser.Active, actualUser.Active)
		}
	}

	// 验证切片不是 nil
	if config.Users == nil {
		t.Error("Users 切片不应该是 nil")
	}

	t.Log("✓ 结构体切片字段解析测试通过")
}

// TestSliceOfStructPointerParsing 测试场景5：结构体的字段是切片且切片的元素是结构体的指针类型的情况下各种类型的字段是否解析正常
func TestSliceOfStructPointerParsing(t *testing.T) {
	// 定义切片元素的结构体指针类型
	type Product struct {
		Name     string  `yaml:"name"`
		Price    float64 `yaml:"price"`
		Quantity int     `yaml:"quantity"`
		InStock  bool    `yaml:"in_stock"`
	}

	type Config struct {
		Products []*Product `yaml:"products"`
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建测试用的 YAML 配置文件，包含结构体指针切片
	yamlContent := `products:
  - name: "Laptop"
    price: 999.99
    quantity: 10
    in_stock: true
  - name: "Mouse"
    price: 29.99
    quantity: 50
    in_stock: true
  - name: "Keyboard"
    price: 79.99
    quantity: 0
    in_stock: false
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证切片长度
	expectedProducts := []*Product{
		{Name: "Laptop", Price: 999.99, Quantity: 10, InStock: true},
		{Name: "Mouse", Price: 29.99, Quantity: 50, InStock: true},
		{Name: "Keyboard", Price: 79.99, Quantity: 0, InStock: false},
	}

	if len(config.Products) != len(expectedProducts) {
		t.Fatalf("Products 切片长度错误: 期望 %d, 实际 %d", len(expectedProducts), len(config.Products))
	}

	// 验证每个结构体指针元素的字段
	for i, expectedProduct := range expectedProducts {
		actualProduct := config.Products[i]
		if actualProduct == nil {
			t.Fatalf("Products[%d] 不应该是 nil", i)
		}
		if actualProduct.Name != expectedProduct.Name {
			t.Errorf("Products[%d].Name 解析错误: 期望 %q, 实际 %q", i, expectedProduct.Name, actualProduct.Name)
		}
		if actualProduct.Price != expectedProduct.Price {
			t.Errorf("Products[%d].Price 解析错误: 期望 %f, 实际 %f", i, expectedProduct.Price, actualProduct.Price)
		}
		if actualProduct.Quantity != expectedProduct.Quantity {
			t.Errorf("Products[%d].Quantity 解析错误: 期望 %d, 实际 %d", i, expectedProduct.Quantity, actualProduct.Quantity)
		}
		if actualProduct.InStock != expectedProduct.InStock {
			t.Errorf("Products[%d].InStock 解析错误: 期望 %v, 实际 %v", i, expectedProduct.InStock, actualProduct.InStock)
		}
	}

	// 验证切片不是 nil
	if config.Products == nil {
		t.Error("Products 切片不应该是 nil")
	}

	t.Log("✓ 结构体指针切片字段解析测试通过")
}

// TestStructPointerFieldParsing 测试场景6：结构体的字段是结构体指针类型且不为nil
func TestStructPointerFieldParsing(t *testing.T) {
	// 定义结构体指针类型
	type DatabaseConfig struct {
		Host     string  `yaml:"host"`
		Port     int     `yaml:"port"`
		Username string  `yaml:"username"`
		Password string  `yaml:"password"`
		Timeout  float64 `yaml:"timeout"`
		Enabled  bool    `yaml:"enabled"`
	}

	type Config struct {
		AppName  string          `yaml:"app_name"`
		Database *DatabaseConfig `yaml:"database"`
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建测试用的 YAML 配置文件，包含结构体指针字段（不为nil）
	yamlContent := `app_name: "myapp"
database:
  host: "localhost"
  port: 5432
  username: "admin"
  password: "secret123"
  timeout: 30.5
  enabled: true
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证顶层字段
	if config.AppName != "myapp" {
		t.Errorf("AppName 解析错误: 期望 %q, 实际 %q", "myapp", config.AppName)
	}

	// 验证结构体指针字段不为 nil（根据 README 要求：对于结构体指针类型，代码会始终保证不是nil）
	if config.Database == nil {
		t.Fatal("Database 结构体指针不应该是 nil")
	}

	// 验证结构体指针指向的结构体字段
	if config.Database.Host != "localhost" {
		t.Errorf("Database.Host 解析错误: 期望 %q, 实际 %q", "localhost", config.Database.Host)
	}

	if config.Database.Port != 5432 {
		t.Errorf("Database.Port 解析错误: 期望 %d, 实际 %d", 5432, config.Database.Port)
	}

	if config.Database.Username != "admin" {
		t.Errorf("Database.Username 解析错误: 期望 %q, 实际 %q", "admin", config.Database.Username)
	}

	if config.Database.Password != "secret123" {
		t.Errorf("Database.Password 解析错误: 期望 %q, 实际 %q", "secret123", config.Database.Password)
	}

	if config.Database.Timeout != 30.5 {
		t.Errorf("Database.Timeout 解析错误: 期望 %f, 实际 %f", 30.5, config.Database.Timeout)
	}

	if config.Database.Enabled != true {
		t.Errorf("Database.Enabled 解析错误: 期望 %v, 实际 %v", true, config.Database.Enabled)
	}

	t.Log("✓ 结构体指针字段（不为nil）解析测试通过")
}

// TestStructPointerFieldNilParsing 测试场景7：结构体的字段是结构体指针类型且为nil
func TestStructPointerFieldNilParsing(t *testing.T) {
	// 定义结构体指针类型
	type DatabaseConfig struct {
		Host     string  `yaml:"host"`
		Port     int     `yaml:"port"`
		Username string  `yaml:"username"`
	}

	type Config struct {
		AppName  string          `yaml:"app_name"`
		Database *DatabaseConfig `yaml:"database"`
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建测试用的 YAML 配置文件，不包含 database 字段（应该为 nil，但代码会保证不是 nil）
	yamlContent := `app_name: "myapp"
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证顶层字段
	if config.AppName != "myapp" {
		t.Errorf("AppName 解析错误: 期望 %q, 实际 %q", "myapp", config.AppName)
	}

	// 验证结构体指针字段不为 nil（根据 README 要求：对于结构体指针类型，代码会始终保证不是nil）
	// 即使 YAML 中没有该字段，代码也应该创建一个空的结构体实例
	if config.Database == nil {
		t.Fatal("Database 结构体指针不应该是 nil，代码应该保证结构体指针不会是 nil")
	}

	// 验证结构体指针指向的结构体字段为零值
	if config.Database.Host != "" {
		t.Errorf("Database.Host 应该是零值，实际 %q", config.Database.Host)
	}

	if config.Database.Port != 0 {
		t.Errorf("Database.Port 应该是零值，实际 %d", config.Database.Port)
	}

	if config.Database.Username != "" {
		t.Errorf("Database.Username 应该是零值，实际 %q", config.Database.Username)
	}

	t.Log("✓ 结构体指针字段（为nil时自动创建）解析测试通过")
}

// TestSliceFieldNotNullParsing 测试场景8：结构体的字段是切片类型且不为nil
func TestSliceFieldNotNullParsing(t *testing.T) {
	type Config struct {
		AppName string   `yaml:"app_name"`
		Tags    []string `yaml:"tags"`
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建测试用的 YAML 配置文件，包含切片字段（不为nil）
	yamlContent := `app_name: "myapp"
tags:
  - "tag1"
  - "tag2"
  - "tag3"
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证顶层字段
	if config.AppName != "myapp" {
		t.Errorf("AppName 解析错误: 期望 %q, 实际 %q", "myapp", config.AppName)
	}

	// 验证切片字段不为 nil（根据 README 要求：对于切片，代码会始终保证不是nil切片）
	if config.Tags == nil {
		t.Fatal("Tags 切片不应该是 nil")
	}

	// 验证切片内容
	expectedTags := []string{"tag1", "tag2", "tag3"}
	if len(config.Tags) != len(expectedTags) {
		t.Fatalf("Tags 切片长度错误: 期望 %d, 实际 %d", len(expectedTags), len(config.Tags))
	}

	for i, expectedTag := range expectedTags {
		if config.Tags[i] != expectedTag {
			t.Errorf("Tags[%d] 解析错误: 期望 %q, 实际 %q", i, expectedTag, config.Tags[i])
		}
	}

	t.Log("✓ 切片字段（不为nil）解析测试通过")
}

// TestSliceFieldNilParsing 测试场景9：结构体的字段是切片类型且为nil
func TestSliceFieldNilParsing(t *testing.T) {
	type Config struct {
		AppName string   `yaml:"app_name"`
		Tags    []string `yaml:"tags"`
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建测试用的 YAML 配置文件，不包含 tags 字段（应该为 nil，但代码会保证不是 nil）
	yamlContent := `app_name: "myapp"
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证顶层字段
	if config.AppName != "myapp" {
		t.Errorf("AppName 解析错误: 期望 %q, 实际 %q", "myapp", config.AppName)
	}

	// 验证切片字段不为 nil（根据 README 要求：对于切片，代码会始终保证不是nil切片）
	// 即使 YAML 中没有该字段，代码也应该创建一个空切片
	if config.Tags == nil {
		t.Fatal("Tags 切片不应该是 nil，代码应该保证切片不会是 nil")
	}

	// 验证切片为空切片（长度为0）
	if len(config.Tags) != 0 {
		t.Errorf("Tags 应该是空切片，实际长度 %d", len(config.Tags))
	}

	t.Log("✓ 切片字段（为nil时自动创建空切片）解析测试通过")
}

// ========== 配置文件默认值测试 ==========

// TestBasicTypesDefaultValues 测试场景1（默认值）：基本类型字段的默认值是否正常赋值
func TestBasicTypesDefaultValues(t *testing.T) {
	type Config struct {
		StringField  string  `yaml:"string_field" default:"default string"`
		IntField     int     `yaml:"int_field" default:"42"`
		Int64Field   int64   `yaml:"int64_field" default:"100"`
		UintField    uint    `yaml:"uint_field" default:"200"`
		Float64Field float64 `yaml:"float64_field" default:"3.14"`
		BoolField    bool    `yaml:"bool_field"` // bool 不支持默认值
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建空的 YAML 配置文件（不包含任何字段，测试默认值）
	yamlContent := ``
	// 或者只包含部分字段
	yamlContent = `int_field: 99
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证默认值是否正确设置
	if config.StringField != "default string" {
		t.Errorf("StringField 默认值错误: 期望 %q, 实际 %q", "default string", config.StringField)
	}

	// int_field 在配置文件中指定了值，应该使用配置文件的值，而不是默认值
	if config.IntField != 99 {
		t.Errorf("IntField 应该使用配置文件的值: 期望 %d, 实际 %d", 99, config.IntField)
	}

	// int64_field 没有在配置文件中，应该使用默认值
	if config.Int64Field != 100 {
		t.Errorf("Int64Field 默认值错误: 期望 %d, 实际 %d", 100, config.Int64Field)
	}

	if config.UintField != 200 {
		t.Errorf("UintField 默认值错误: 期望 %d, 实际 %d", 200, config.UintField)
	}

	if config.Float64Field != 3.14 {
		t.Errorf("Float64Field 默认值错误: 期望 %f, 实际 %f", 3.14, config.Float64Field)
	}

	// bool 不支持默认值，应该保持零值 false
	if config.BoolField != false {
		t.Errorf("BoolField 应该不支持默认值，保持零值 false，实际 %v", config.BoolField)
	}

	t.Log("✓ 基本类型字段默认值测试通过")
}

// TestNestedStructDefaultValues 测试场景2（默认值）：嵌套结构体中字段的默认值是否正常赋值
func TestNestedStructDefaultValues(t *testing.T) {
	type DatabaseConfig struct {
		Host     string  `yaml:"host" default:"localhost"`
		Port     int     `yaml:"port" default:"5432"`
		Username string  `yaml:"username" default:"admin"`
		Timeout  float64 `yaml:"timeout" default:"30.0"`
	}

	type Config struct {
		AppName  string         `yaml:"app_name" default:"myapp"`
		Database DatabaseConfig `yaml:"database"` // 结构体本身不支持默认值
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建空的 YAML 配置文件
	yamlContent := ``
	// 或者只包含部分字段
	yamlContent = `app_name: "customapp"
database:
  port: 3306
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证顶层字段的默认值
	if config.AppName != "customapp" {
		t.Errorf("AppName 应该使用配置文件的值: 期望 %q, 实际 %q", "customapp", config.AppName)
	}

	// 验证嵌套结构体中字段的默认值
	if config.Database.Host != "localhost" {
		t.Errorf("Database.Host 默认值错误: 期望 %q, 实际 %q", "localhost", config.Database.Host)
	}

	// port 在配置文件中指定了值，应该使用配置文件的值
	if config.Database.Port != 3306 {
		t.Errorf("Database.Port 应该使用配置文件的值: 期望 %d, 实际 %d", 3306, config.Database.Port)
	}

	if config.Database.Username != "admin" {
		t.Errorf("Database.Username 默认值错误: 期望 %q, 实际 %q", "admin", config.Database.Username)
	}

	if config.Database.Timeout != 30.0 {
		t.Errorf("Database.Timeout 默认值错误: 期望 %f, 实际 %f", 30.0, config.Database.Timeout)
	}

	t.Log("✓ 嵌套结构体字段默认值测试通过")
}

// TestStructPointerDefaultValues 测试场景6（默认值）：结构体指针指向的结构体中字段的默认值是否正常赋值
func TestStructPointerDefaultValues(t *testing.T) {
	type DatabaseConfig struct {
		Host     string  `yaml:"host" default:"localhost"`
		Port     int     `yaml:"port" default:"5432"`
		Username string  `yaml:"username" default:"admin"`
	}

	type Config struct {
		AppName  string          `yaml:"app_name" default:"myapp"`
		Database *DatabaseConfig `yaml:"database"` // 结构体指针本身不支持默认值
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建空的 YAML 配置文件
	yamlContent := ``
	// 或者只包含部分字段
	yamlContent = `database:
  port: 3306
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证结构体指针不为 nil
	if config.Database == nil {
		t.Fatal("Database 结构体指针不应该是 nil")
	}

	// 验证结构体指针指向的结构体中字段的默认值
	if config.Database.Host != "localhost" {
		t.Errorf("Database.Host 默认值错误: 期望 %q, 实际 %q", "localhost", config.Database.Host)
	}

	// port 在配置文件中指定了值，应该使用配置文件的值
	if config.Database.Port != 3306 {
		t.Errorf("Database.Port 应该使用配置文件的值: 期望 %d, 实际 %d", 3306, config.Database.Port)
	}

	if config.Database.Username != "admin" {
		t.Errorf("Database.Username 默认值错误: 期望 %q, 实际 %q", "admin", config.Database.Username)
	}

	t.Log("✓ 结构体指针字段默认值测试通过")
}

// TestSliceOfStructDefaultValues 测试场景4（默认值）：切片中结构体元素的字段默认值是否正常赋值
func TestSliceOfStructDefaultValues(t *testing.T) {
	type User struct {
		Name  string  `yaml:"name" default:"Unknown"`
		Age   int     `yaml:"age" default:"18"`
		Email string  `yaml:"email" default:"user@example.com"`
		Score float64 `yaml:"score" default:"0.0"`
	}

	type Config struct {
		Users []User `yaml:"users"` // 切片本身不支持默认值
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建 YAML 配置文件，部分字段缺失
	yamlContent := `users:
  - name: "Alice"
    age: 25
  - name: "Bob"
    score: 87.0
  - age: 30
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证切片长度
	if len(config.Users) != 3 {
		t.Fatalf("Users 切片长度错误: 期望 %d, 实际 %d", 3, len(config.Users))
	}

	// 验证第一个用户：name 和 age 有值，email 和 score 应该使用默认值
	if config.Users[0].Name != "Alice" {
		t.Errorf("Users[0].Name 错误: 期望 %q, 实际 %q", "Alice", config.Users[0].Name)
	}
	if config.Users[0].Age != 25 {
		t.Errorf("Users[0].Age 错误: 期望 %d, 实际 %d", 25, config.Users[0].Age)
	}
	if config.Users[0].Email != "user@example.com" {
		t.Errorf("Users[0].Email 默认值错误: 期望 %q, 实际 %q", "user@example.com", config.Users[0].Email)
	}
	if config.Users[0].Score != 0.0 {
		t.Errorf("Users[0].Score 默认值错误: 期望 %f, 实际 %f", 0.0, config.Users[0].Score)
	}

	// 验证第二个用户：name 和 score 有值，age 和 email 应该使用默认值
	if config.Users[1].Name != "Bob" {
		t.Errorf("Users[1].Name 错误: 期望 %q, 实际 %q", "Bob", config.Users[1].Name)
	}
	if config.Users[1].Age != 18 {
		t.Errorf("Users[1].Age 默认值错误: 期望 %d, 实际 %d", 18, config.Users[1].Age)
	}
	if config.Users[1].Email != "user@example.com" {
		t.Errorf("Users[1].Email 默认值错误: 期望 %q, 实际 %q", "user@example.com", config.Users[1].Email)
	}
	if config.Users[1].Score != 87.0 {
		t.Errorf("Users[1].Score 错误: 期望 %f, 实际 %f", 87.0, config.Users[1].Score)
	}

	// 验证第三个用户：只有 age 有值，其他应该使用默认值
	if config.Users[2].Name != "Unknown" {
		t.Errorf("Users[2].Name 默认值错误: 期望 %q, 实际 %q", "Unknown", config.Users[2].Name)
	}
	if config.Users[2].Age != 30 {
		t.Errorf("Users[2].Age 错误: 期望 %d, 实际 %d", 30, config.Users[2].Age)
	}
	if config.Users[2].Email != "user@example.com" {
		t.Errorf("Users[2].Email 默认值错误: 期望 %q, 实际 %q", "user@example.com", config.Users[2].Email)
	}
	if config.Users[2].Score != 0.0 {
		t.Errorf("Users[2].Score 默认值错误: 期望 %f, 实际 %f", 0.0, config.Users[2].Score)
	}

	t.Log("✓ 切片中结构体元素字段默认值测试通过")
}

// TestSliceOfStructPointerDefaultValues 测试场景5（默认值）：切片中结构体指针元素的字段默认值是否正常赋值
func TestSliceOfStructPointerDefaultValues(t *testing.T) {
	type Product struct {
		Name     string  `yaml:"name" default:"Unknown Product"`
		Price    float64 `yaml:"price" default:"0.0"`
		Quantity int     `yaml:"quantity" default:"0"`
	}

	type Config struct {
		Products []*Product `yaml:"products"` // 切片本身不支持默认值
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建 YAML 配置文件，部分字段缺失
	yamlContent := `products:
  - name: "Laptop"
    price: 999.99
  - quantity: 50
  - name: "Mouse"
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证切片长度
	if len(config.Products) != 3 {
		t.Fatalf("Products 切片长度错误: 期望 %d, 实际 %d", 3, len(config.Products))
	}

	// 验证第一个产品：name 和 price 有值，quantity 应该使用默认值
	if config.Products[0] == nil {
		t.Fatal("Products[0] 不应该是 nil")
	}
	if config.Products[0].Name != "Laptop" {
		t.Errorf("Products[0].Name 错误: 期望 %q, 实际 %q", "Laptop", config.Products[0].Name)
	}
	if config.Products[0].Price != 999.99 {
		t.Errorf("Products[0].Price 错误: 期望 %f, 实际 %f", 999.99, config.Products[0].Price)
	}
	if config.Products[0].Quantity != 0 {
		t.Errorf("Products[0].Quantity 默认值错误: 期望 %d, 实际 %d", 0, config.Products[0].Quantity)
	}

	// 验证第二个产品：只有 quantity 有值，其他应该使用默认值
	if config.Products[1] == nil {
		t.Fatal("Products[1] 不应该是 nil")
	}
	if config.Products[1].Name != "Unknown Product" {
		t.Errorf("Products[1].Name 默认值错误: 期望 %q, 实际 %q", "Unknown Product", config.Products[1].Name)
	}
	if config.Products[1].Price != 0.0 {
		t.Errorf("Products[1].Price 默认值错误: 期望 %f, 实际 %f", 0.0, config.Products[1].Price)
	}
	if config.Products[1].Quantity != 50 {
		t.Errorf("Products[1].Quantity 错误: 期望 %d, 实际 %d", 50, config.Products[1].Quantity)
	}

	// 验证第三个产品：只有 name 有值，其他应该使用默认值
	if config.Products[2] == nil {
		t.Fatal("Products[2] 不应该是 nil")
	}
	if config.Products[2].Name != "Mouse" {
		t.Errorf("Products[2].Name 错误: 期望 %q, 实际 %q", "Mouse", config.Products[2].Name)
	}
	if config.Products[2].Price != 0.0 {
		t.Errorf("Products[2].Price 默认值错误: 期望 %f, 实际 %f", 0.0, config.Products[2].Price)
	}
	if config.Products[2].Quantity != 0 {
		t.Errorf("Products[2].Quantity 默认值错误: 期望 %d, 实际 %d", 0, config.Products[2].Quantity)
	}

	t.Log("✓ 切片中结构体指针元素字段默认值测试通过")
}

// ========== 其他测试 ==========

// TestUnsupportedDefaultValues 测试：bool、结构体、结构体指针类型、切片不支持默认值，但其内部不受影响
func TestUnsupportedDefaultValues(t *testing.T) {
	type NestedStruct struct {
		Field string `yaml:"field" default:"nested default"`
	}

	type Config struct {
		BoolField          bool           `yaml:"bool_field" default:"true"`           // bool 不支持默认值
		StructField        NestedStruct   `yaml:"struct_field" default:"ignored"`     // 结构体不支持默认值
		StructPointerField *NestedStruct  `yaml:"struct_pointer_field" default:"nil"` // 结构体指针不支持默认值
		SliceField         []string       `yaml:"slice_field" default:"ignored"`      // 切片不支持默认值
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建空的 YAML 配置文件
	yamlContent := `struct_field:
  field: "custom"
struct_pointer_field:
  field: "pointer custom"
slice_field:
  - "item1"
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证 bool 不支持默认值，应该保持零值 false
	if config.BoolField != false {
		t.Errorf("BoolField 应该不支持默认值，保持零值 false，实际 %v", config.BoolField)
	}

	// 验证结构体不支持默认值，但内部字段支持默认值
	// struct_field 在配置文件中指定了 field，应该使用配置文件的值
	if config.StructField.Field != "custom" {
		t.Errorf("StructField.Field 应该使用配置文件的值: 期望 %q, 实际 %q", "custom", config.StructField.Field)
	}

	// 验证结构体指针不支持默认值，但内部字段支持默认值
	if config.StructPointerField == nil {
		t.Fatal("StructPointerField 不应该是 nil")
	}
	if config.StructPointerField.Field != "pointer custom" {
		t.Errorf("StructPointerField.Field 应该使用配置文件的值: 期望 %q, 实际 %q", "pointer custom", config.StructPointerField.Field)
	}

	// 验证切片不支持默认值，但切片本身不会是 nil
	if config.SliceField == nil {
		t.Fatal("SliceField 不应该是 nil")
	}
	if len(config.SliceField) != 1 || config.SliceField[0] != "item1" {
		t.Errorf("SliceField 应该使用配置文件的值: 期望 [%q], 实际 %v", "item1", config.SliceField)
	}

	// 测试结构体内部字段的默认值（当配置文件中没有该字段时）
	type Config2 struct {
		StructField NestedStruct `yaml:"struct_field"`
	}

	tmpDir2 := t.TempDir()
	configFile2 := filepath.Join(tmpDir2, "test_config2.yaml")
	yamlContent2 := `struct_field: {}
`

	if err := os.WriteFile(configFile2, []byte(yamlContent2), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	loader2 := NewYamlLoader(configFile2, "")
	config2 := &Config2{}

	if err := loader2.Load(config2); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证结构体内部字段的默认值生效
	if config2.StructField.Field != "nested default" {
		t.Errorf("StructField.Field 默认值错误: 期望 %q, 实际 %q", "nested default", config2.StructField.Field)
	}

	t.Log("✓ 不支持默认值的类型测试通过")
}

// TestSliceElementNilPointerSkip 测试：切片的元素是结构体指针类型且元素是nil时，该元素跳过任何处理
func TestSliceElementNilPointerSkip(t *testing.T) {
	type Product struct {
		Name     string  `yaml:"name" default:"Unknown"`
		Price    float64 `yaml:"price" default:"0.0"`
		Quantity int     `yaml:"quantity" default:"0"`
	}

	type Config struct {
		Products []*Product `yaml:"products"`
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建 YAML 配置文件，包含 nil 指针元素
	// 注意：YAML 中无法直接表示 nil，但可以通过不完整的数据结构来模拟
	// 实际上，如果 YAML 中某个元素是 null，YAML 解析器会将其解析为 nil
	yamlContent := `products:
  - name: "Laptop"
    price: 999.99
  - null
  - name: "Mouse"
    price: 29.99
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证切片长度
	if len(config.Products) != 3 {
		t.Fatalf("Products 切片长度错误: 期望 %d, 实际 %d", 3, len(config.Products))
	}

	// 验证第一个产品（不为nil）
	if config.Products[0] == nil {
		t.Fatal("Products[0] 不应该是 nil")
	}
	if config.Products[0].Name != "Laptop" {
		t.Errorf("Products[0].Name 错误: 期望 %q, 实际 %q", "Laptop", config.Products[0].Name)
	}

	// 验证第二个产品（为nil，应该跳过处理，保持nil）
	if config.Products[1] != nil {
		t.Error("Products[1] 应该是 nil，nil 指针元素应该跳过处理")
	}

	// 验证第三个产品（不为nil）
	if config.Products[2] == nil {
		t.Fatal("Products[2] 不应该是 nil")
	}
	if config.Products[2].Name != "Mouse" {
		t.Errorf("Products[2].Name 错误: 期望 %q, 实际 %q", "Mouse", config.Products[2].Name)
	}

	t.Log("✓ 切片中nil指针元素跳过处理测试通过")
}

// TestStructPointerFieldNilDefaultValues 测试场景7（默认值）：结构体指针字段为nil时，内部字段的默认值是否正常赋值
func TestStructPointerFieldNilDefaultValues(t *testing.T) {
	type DatabaseConfig struct {
		Host     string  `yaml:"host" default:"localhost"`
		Port     int     `yaml:"port" default:"5432"`
		Username string  `yaml:"username" default:"admin"`
		Timeout  float64 `yaml:"timeout" default:"30.0"`
	}

	type Config struct {
		AppName  string          `yaml:"app_name" default:"myapp"`
		Database *DatabaseConfig `yaml:"database"` // 结构体指针本身不支持默认值
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建空的 YAML 配置文件（不包含 database 字段）
	yamlContent := `app_name: "customapp"
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证顶层字段
	if config.AppName != "customapp" {
		t.Errorf("AppName 应该使用配置文件的值: 期望 %q, 实际 %q", "customapp", config.AppName)
	}

	// 验证结构体指针不为 nil（代码会自动创建）
	if config.Database == nil {
		t.Fatal("Database 结构体指针不应该是 nil，代码应该自动创建")
	}

	// 验证结构体指针指向的结构体中字段的默认值
	if config.Database.Host != "localhost" {
		t.Errorf("Database.Host 默认值错误: 期望 %q, 实际 %q", "localhost", config.Database.Host)
	}

	if config.Database.Port != 5432 {
		t.Errorf("Database.Port 默认值错误: 期望 %d, 实际 %d", 5432, config.Database.Port)
	}

	if config.Database.Username != "admin" {
		t.Errorf("Database.Username 默认值错误: 期望 %q, 实际 %q", "admin", config.Database.Username)
	}

	if config.Database.Timeout != 30.0 {
		t.Errorf("Database.Timeout 默认值错误: 期望 %f, 实际 %f", 30.0, config.Database.Timeout)
	}

	t.Log("✓ 结构体指针字段（为nil时）默认值测试通过")
}

// TestSliceFieldNilDefaultValues 测试场景9（默认值）：切片字段为nil时，切片中结构体元素的字段默认值是否正常赋值
func TestSliceFieldNilDefaultValues(t *testing.T) {
	type User struct {
		Name  string  `yaml:"name" default:"Unknown"`
		Age   int     `yaml:"age" default:"18"`
		Email string  `yaml:"email" default:"user@example.com"`
		Score float64 `yaml:"score" default:"0.0"`
	}

	type Config struct {
		AppName string  `yaml:"app_name" default:"myapp"`
		Users   []User  `yaml:"users"` // 切片本身不支持默认值
	}

	// 创建临时测试目录
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yaml")

	// 创建空的 YAML 配置文件（不包含 users 字段，切片为nil，代码会创建空切片）
	yamlContent := `app_name: "customapp"
`

	if err := os.WriteFile(configFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	// 创建配置加载器
	loader := NewYamlLoader(configFile, "")

	// 创建配置实例
	config := &Config{}

	// 加载配置
	if err := loader.Load(config); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证顶层字段
	if config.AppName != "customapp" {
		t.Errorf("AppName 应该使用配置文件的值: 期望 %q, 实际 %q", "customapp", config.AppName)
	}

	// 验证切片不为 nil（代码会自动创建空切片）
	if config.Users == nil {
		t.Fatal("Users 切片不应该是 nil，代码应该自动创建空切片")
	}

	// 验证切片为空切片（长度为0）
	if len(config.Users) != 0 {
		t.Errorf("Users 应该是空切片，实际长度 %d", len(config.Users))
	}

	// 现在测试：如果切片中有元素，元素字段的默认值应该生效
	type Config2 struct {
		Users []User `yaml:"users"`
	}

	tmpDir2 := t.TempDir()
	configFile2 := filepath.Join(tmpDir2, "test_config2.yaml")
	yamlContent2 := `users:
  - {}
  - name: "Alice"
`

	if err := os.WriteFile(configFile2, []byte(yamlContent2), 0644); err != nil {
		t.Fatalf("创建测试配置文件失败: %v", err)
	}

	loader2 := NewYamlLoader(configFile2, "")
	config2 := &Config2{}

	if err := loader2.Load(config2); err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证第一个用户：所有字段都应该使用默认值
	if config2.Users[0].Name != "Unknown" {
		t.Errorf("Users[0].Name 默认值错误: 期望 %q, 实际 %q", "Unknown", config2.Users[0].Name)
	}
	if config2.Users[0].Age != 18 {
		t.Errorf("Users[0].Age 默认值错误: 期望 %d, 实际 %d", 18, config2.Users[0].Age)
	}
	if config2.Users[0].Email != "user@example.com" {
		t.Errorf("Users[0].Email 默认值错误: 期望 %q, 实际 %q", "user@example.com", config2.Users[0].Email)
	}
	if config2.Users[0].Score != 0.0 {
		t.Errorf("Users[0].Score 默认值错误: 期望 %f, 实际 %f", 0.0, config2.Users[0].Score)
	}

	// 验证第二个用户：name 有值，其他字段使用默认值
	if config2.Users[1].Name != "Alice" {
		t.Errorf("Users[1].Name 错误: 期望 %q, 实际 %q", "Alice", config2.Users[1].Name)
	}
	if config2.Users[1].Age != 18 {
		t.Errorf("Users[1].Age 默认值错误: 期望 %d, 实际 %d", 18, config2.Users[1].Age)
	}

	t.Log("✓ 切片字段（为nil时）默认值测试通过")
}
