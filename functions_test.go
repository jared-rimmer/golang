package main

import "testing"

func TestSetAge(t *testing.T) {
	test_person := Person{"Jared", 96, []int{10, 20, 30}, "Golang"}
	expected := 90
	test_person.setAge(40)
	if test_person.age != 40 {
		t.Errorf("test_person.setAge(40) = %d; want %d", test_person.age, expected)
	}
}

func TestSetSubject(t *testing.T) {
	test_person := Person{"Jared", 96, []int{10, 20, 30}, "Python"}
	expected := "Golang"
	test_person.setSubject("Golang")
	if test_person.subject != "Golang" {
		t.Errorf("test_person.setSubect('Golang') = %s; want %s", test_person.subject, expected)
	}
}

func TestCalculateAverageGrade(t *testing.T) {
	test_person := Person{"Jared", 96, []int{10, 20, 30}, "Python"}
	expected := 20
	result := test_person.calculateAverageGrade()
	if result != float32(expected) {
		t.Errorf("test_person.TestCalculateAverageGrade = %v; want %v", result, expected)
	}
}
