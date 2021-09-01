package file

import (
	"fmt"
	"os"

	json "github.com/json-iterator/go"
)

const (
	ProjectPath = "" // 项目文件存放位置
)

// SaveAsJson 将结构体以JSON的形式存到指定文件中
func SaveAsJson(filename string, obj interface{}) error {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("SaveAsJson: %w", err)
	}

	err = os.WriteFile(filename, jsonBytes, 0774)
	if err != nil {
		return fmt.Errorf("SaveAsJson: %w", err)
	}

	return nil
}

// ReadFromJson 读取指定文件，并解析JSON到指定的结构体中
// 如果没有文件，则会返回空值，并创建文件
func ReadFromJson(filename string, obj interface{}) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("ReadFile: %w", err)
	}

	err = json.Unmarshal(content, obj)
	if err != nil {
		return fmt.Errorf("ReadFile: %w", err)
	}
	return nil
}
