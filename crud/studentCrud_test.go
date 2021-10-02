package crud

import (
	"IceBreaking/model"
	"fmt"
	"testing"
)

func TestAddStudent(t *testing.T) {
	// 测试 BeforeCreate 接口
	err := AddStudent(&model.Student{StaffId: "123456", Name: "testUser1"})
	if err != nil {
		fmt.Println(err)
	}
}
