// Reflection: introspeksi dan manipulasi runtime types dan values
package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Person struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"min=0,max=150"`
	Email   string `json:"email" validate:"email"`
	Address Address
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func inspectType(v interface{}) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	
	fmt.Printf("Type: %s, Kind: %s\n", t, t.Kind())
	fmt.Printf("Value: %v, Type: %T\n", val.Interface(), val.Interface())
	
	if t.Kind() == reflect.Ptr {
		fmt.Printf("Pointer to: %s\n", t.Elem())
		if !val.IsNil() {
			val = val.Elem()
			t = t.Elem()
		}
	}
	
	if t.Kind() == reflect.Struct {
		fmt.Printf("Struct has %d fields:\n", t.NumField())
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := val.Field(i)
			fmt.Printf("  Field %d: %s (%s) = %v\n", i, field.Name, field.Type, fieldValue.Interface())
			
			if tag := field.Tag.Get("json"); tag != "" {
				fmt.Printf("    JSON tag: %s\n", tag)
			}
			if tag := field.Tag.Get("validate"); tag != "" {
				fmt.Printf("    Validate tag: %s\n", tag)
			}
		}
	}
}

func modifyStruct(v interface{}) {
	val := reflect.ValueOf(v)
	
	if val.Kind() != reflect.Ptr {
		fmt.Println("Cannot modify non-pointer value")
		return
	}
	
	val = val.Elem()
	if val.Kind() != reflect.Struct {
		fmt.Println("Value is not a struct")
		return
	}
	
	nameField := val.FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() {
		if nameField.Kind() == reflect.String {
			nameField.SetString("Modified Name")
			fmt.Println("Name field modified")
		}
	}
	
	ageField := val.FieldByName("Age")
	if ageField.IsValid() && ageField.CanSet() {
		if ageField.Kind() == reflect.Int {
			ageField.SetInt(99)
			fmt.Println("Age field modified")
		}
	}
}

func callMethod(v interface{}, methodName string, args ...interface{}) {
	val := reflect.ValueOf(v)
	method := val.MethodByName(methodName)
	
	if !method.IsValid() {
		fmt.Printf("Method %s not found\n", methodName)
		return
	}
	
	argValues := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}
	
	results := method.Call(argValues)
	
	fmt.Printf("Method %s called with results:\n", methodName)
	for i, result := range results {
		fmt.Printf("  Result %d: %v\n", i, result.Interface())
	}
}

type Calculator struct {
	value float64
}

func (c *Calculator) Add(x float64) float64 {
	c.value += x
	return c.value
}

func (c *Calculator) Multiply(x float64) float64 {
	c.value *= x
	return c.value
}

func (c Calculator) GetValue() float64 {
	return c.value
}

func createStruct(typeName string) interface{} {
	switch typeName {
	case "Person":
		return &Person{}
	case "Address":
		return &Address{}
	default:
		return nil
	}
}

func validateStruct(v interface{}) []string {
	var errors []string
	
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	
	t := val.Type()
	
	for i := 0; i < val.NumField(); i++ {
		field := t.Field(i)
		fieldValue := val.Field(i)
		validateTag := field.Tag.Get("validate")
		
		if validateTag == "" {
			continue
		}
		
		rules := strings.Split(validateTag, ",")
		for _, rule := range rules {
			if rule == "required" {
				if isZeroValue(fieldValue) {
					errors = append(errors, fmt.Sprintf("Field %s is required", field.Name))
				}
			} else if strings.HasPrefix(rule, "min=") {
				min, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))
				if fieldValue.Kind() == reflect.Int && int(fieldValue.Int()) < min {
					errors = append(errors, fmt.Sprintf("Field %s must be at least %d", field.Name, min))
				}
			} else if strings.HasPrefix(rule, "max=") {
				max, _ := strconv.Atoi(strings.TrimPrefix(rule, "max="))
				if fieldValue.Kind() == reflect.Int && int(fieldValue.Int()) > max {
					errors = append(errors, fmt.Sprintf("Field %s must be at most %d", field.Name, max))
				}
			}
		}
	}
	
	return errors
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	}
}

func deepCopy(src interface{}) interface{} {
	srcVal := reflect.ValueOf(src)
	return deepCopyValue(srcVal).Interface()
}

func deepCopyValue(src reflect.Value) reflect.Value {
	switch src.Kind() {
	case reflect.Ptr:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		dst := reflect.New(src.Type().Elem())
		dst.Elem().Set(deepCopyValue(src.Elem()))
		return dst
	case reflect.Struct:
		dst := reflect.New(src.Type()).Elem()
		for i := 0; i < src.NumField(); i++ {
			dst.Field(i).Set(deepCopyValue(src.Field(i)))
		}
		return dst
	case reflect.Slice:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		dst := reflect.MakeSlice(src.Type(), src.Len(), src.Cap())
		for i := 0; i < src.Len(); i++ {
			dst.Index(i).Set(deepCopyValue(src.Index(i)))
		}
		return dst
	case reflect.Map:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		dst := reflect.MakeMap(src.Type())
		for _, key := range src.MapKeys() {
			dst.SetMapIndex(key, deepCopyValue(src.MapIndex(key)))
		}
		return dst
	default:
		return src
	}
}

func main() {
	fmt.Println("=== BASIC REFLECTION ===")
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
	}
	
	inspectType(person)
	
	fmt.Println("\n=== REFLECTION ON DIFFERENT TYPES ===")
	inspectType(42)
	inspectType("hello")
	inspectType([]int{1, 2, 3})
	inspectType(&person)
	
	fmt.Println("\n=== MODIFYING VALUES ===")
	fmt.Printf("Before modification: %+v\n", person)
	modifyStruct(&person)
	fmt.Printf("After modification: %+v\n", person)
	
	fmt.Println("\n=== METHOD CALLING ===")
	calc := &Calculator{value: 10}
	fmt.Printf("Initial calculator: %+v\n", calc)
	
	callMethod(calc, "Add", 5.0)
	callMethod(calc, "Multiply", 2.0)
	callMethod(calc, "GetValue")
	
	fmt.Println("\n=== DYNAMIC STRUCT CREATION ===")
	dynamicPerson := createStruct("Person")
	if dynamicPerson != nil {
		fmt.Printf("Created struct: %T\n", dynamicPerson)
		inspectType(dynamicPerson)
	}
	
	fmt.Println("\n=== STRUCT VALIDATION ===")
	validPerson := Person{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
	}
	
	invalidPerson := Person{
		Name: "",
		Age:  -5,
	}
	
	fmt.Println("Validating valid person:")
	if errors := validateStruct(validPerson); len(errors) == 0 {
		fmt.Println("Valid!")
	} else {
		for _, err := range errors {
			fmt.Printf("  Error: %s\n", err)
		}
	}
	
	fmt.Println("Validating invalid person:")
	if errors := validateStruct(invalidPerson); len(errors) == 0 {
		fmt.Println("Valid!")
	} else {
		for _, err := range errors {
			fmt.Printf("  Error: %s\n", err)
		}
	}
	
	fmt.Println("\n=== TYPE SWITCHING ===")
	values := []interface{}{42, "hello", 3.14, true, []int{1, 2, 3}}
	
	for _, v := range values {
		val := reflect.ValueOf(v)
		switch val.Kind() {
		case reflect.Int:
			fmt.Printf("Integer: %d\n", val.Int())
		case reflect.String:
			fmt.Printf("String: %s\n", val.String())
		case reflect.Float64:
			fmt.Printf("Float: %.2f\n", val.Float())
		case reflect.Bool:
			fmt.Printf("Boolean: %t\n", val.Bool())
		case reflect.Slice:
			fmt.Printf("Slice of %s with length %d\n", val.Type().Elem(), val.Len())
		default:
			fmt.Printf("Unknown type: %s\n", val.Type())
		}
	}
	
	fmt.Println("\n=== DEEP COPY ===")
	original := Person{
		Name: "Original",
		Age:  30,
		Address: Address{
			Street: "Original Street",
			City:   "Original City",
		},
	}
	
	copied := deepCopy(original).(Person)
	copied.Name = "Copied"
	copied.Address.Street = "Copied Street"
	
	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Copied: %+v\n", copied)
}
