package gonfig

import (
	"bytes"
	"fmt"
	//"log"
	"testing"
)

func TestString1(t *testing.T) {
	reader := bytes.NewReader([]byte(`
		{
			"some" : {
				"key" : "abc"
			}
		}
	`))
	config, err := FromJson(reader)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	value, err := config.GetString("some/key", "")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if value != "abc" {
		t.Error(fmt.Sprintf("expected value to be abc but it was %v", value))
		t.Fail()
	}
}

func TestStringDefault1(t *testing.T) {
	reader := bytes.NewReader([]byte(`
		{
			"some" : {
			}
		}
	`))
	config, err := FromJson(reader)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	value, err := config.GetString("what", "monkey")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if value != "monkey" {
		t.Error(fmt.Sprintf("expected value to be monkey but it was %v", value))
		t.Fail()
	}
}

func TestBool1(t *testing.T) {
	reader := bytes.NewReader([]byte(`
		{
			"some" : {
				"key" : true
			}
		}
	`))
	config, err := FromJson(reader)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	value, err := config.GetBool("some/key", false)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if value != true {
		t.Error(fmt.Sprintf("expected value to be true but it was %v", value))
		t.Fail()
	}
}

func TestBoolDefault1(t *testing.T) {
	reader := bytes.NewReader([]byte(`
		{
			"some" : {
			}
		}
	`))
	config, err := FromJson(reader)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	value, err := config.GetBool("some/key/monkey/banana", true)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if value != true {
		t.Error(fmt.Sprintf("expected value to be true but it was %v", value))
		t.Fail()
	}
}

func TestInt1(t *testing.T) {
	reader := bytes.NewReader([]byte(`
		{
			"some" : {
				"key" : {
					"value" : 56
				}
			}
		}
	`))
	config, err := FromJson(reader)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	value, err := config.GetInt("some/key/value", nil)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if value != 56 {
		t.Error(fmt.Sprintf("expected value to be 56 but it was %v", value))
		t.Fail()
	}
}

func TestIntDefault1(t *testing.T) {
	reader := bytes.NewReader([]byte(`
		{
			"some" : {
			}
		}
	`))
	config, err := FromJson(reader)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	value, err := config.GetInt("monkey/banana", 1569)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if value != 1569 {
		t.Error(fmt.Sprintf("expected value to be 1569 but it was %v", value))
		t.Fail()
	}
}

func TestFloat1(t *testing.T) {
	reader := bytes.NewReader([]byte(`
		{
			"some" : {
				"key" : {
					"value" : 56
				},
				"value" : 123.56
			}
		}
	`))
	config, err := FromJson(reader)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	value, err := config.GetFloat("some/value", nil)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if value != 123.56 {
		t.Error(fmt.Sprintf("expected value to be 123.56 but it was %v", value))
		t.Fail()
	}
}

func TestFloatDefault1(t *testing.T) {
	reader := bytes.NewReader([]byte(`
		{
			"some" : {
			}
		}
	`))
	config, err := FromJson(reader)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	value, err := config.GetFloat("monkey/banana", 123.56)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if value != 123.56 {
		t.Error(fmt.Sprintf("expected value to be 123.56 but it was %v", value))
		t.Fail()
	}
}

type TestStruct struct {
	Monkeys    int
	List       []string
	Active     bool
	Percentage float64
}

func TestGetAs(t *testing.T) {
	reader := bytes.NewReader([]byte(`
		{
			"some" : {
				"key" : {
					"monkeys" : 5000,
					"list" : ["one", "two", "three"],
					"active" : true,
					"percentage" : 0.56
				}
			}
		}
	`))
	config, err := FromJson(reader)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	var out TestStruct
	err = config.GetAs("some/key", &out)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if out.Monkeys != 5000 {
		t.Error("expected 5000 monkeys")
		t.Fail()
	}
	if out.Active != true {
		t.Error("expected active to be true")
		t.Fail()
	}
	if out.Percentage != 0.56 {
		t.Error("expected percentage to be 0.56")
		t.Fail()
	}
	if len(out.List) != 3 {
		t.Error("expected a list of 3")
		t.Fail()
	}
}
