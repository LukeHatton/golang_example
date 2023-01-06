package greetings

// 运行测试函数需要导入testing包
// 使用regexp检测函数的输入输出是否与预期相匹配
import (
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
)

// 单元测试函数接受一个指向类型为testing.T的指针参数
// 我们会使用T这个类型的方法来进行单元测试的记录与报告
func TestHelloName(t *testing.T) {
	name := "Tom"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	// 如果测试函数没有按照预期输出，就使用Fatalf来打印一个错误信息，并结束执行
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Tom") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", but got an error`, msg, err)
	}
}

func TestHellos(t *testing.T) {
	assertion := assert.New(t)
	names := []string{"ming", "hong", "xia"}
	_, err := Hellos(names)
	if err != nil {
		log.Fatalln(err)
	}
	names2 := []string{""}
	_, err2 := Hellos(names2)
	assertion.EqualError(err2, "empty name")
}
