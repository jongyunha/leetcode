package easy

import "testing"

func Test_lemonade_test(t *testing.T) {
	bills := []int{5, 5, 5, 10, 20}
	result := lemonadeChange(bills)
	if result != true {
		t.Error("Wrong result")
	}

	bills = []int{5, 5, 10}
	result = lemonadeChange(bills)
	if result != true {
		t.Error("Wrong result")
	}

	bills = []int{10, 10}
	result = lemonadeChange(bills)
	if result != false {
		t.Error("Wrong result")
	}

	bills = []int{5, 5, 10, 10, 20}
	result = lemonadeChange(bills)
	if result != false {
		t.Error("Wrong result")
	}

	bills = []int{5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20}
	result = lemonadeChange(bills)
	if result != true {
		t.Error("Wrong result")
	}

	bills = []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}
	result = lemonadeChange(bills)
	if result != true {
		t.Error("Wrong result")
	}
}
