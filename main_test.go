package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	testValues := []FileStruct{
		{
			Size: 10,
			Name: "Ssdf",
		},
		{
			Size: 100,
			Name: "File",
		},
		{
			Size: 30,
			Name: "File",
		},
		{
			Size: 40,
			Name: "New file",
		},
	}

	expected := []FileStruct{
		{
			Size: 100,
			Name: "File",
		},
		{
			Size: 40,
			Name: "New file",
		},
		{
			Size: 30,
			Name: "File (3)",
		},
		{
			Size: 10,
			Name: "Ssdf",
		},
	}

	makeUniqueFileNames(testValues)

	fmt.Println("testVaules", testValues)
	fmt.Println("expected", expected)

	if !reflect.DeepEqual(testValues, expected) {
		t.Errorf("Expected %+v, but got another %+v", expected, testValues)
	}

}
