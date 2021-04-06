package base

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `mytag:"A"`
	Email string `mytag:"B"`
}

func HandleMyTag() {
	u := User{"bill", "bill@my.com"}
	t := reflect.TypeOf(u)

	for _, fieldName := range []string{"Name", "Email"} {
		field, found := t.FieldByName(fieldName)
		if !found {
			continue
		}
		fmt.Printf("\nField: User.%s\n", fieldName)
		fmt.Printf("\tWhole tag value : %q\n", field.Tag)
		fmt.Printf("\tValue of 'mytag': %q\n", field.Tag.Get("mytag"))
	}
}
