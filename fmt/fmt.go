//// https://blog.gopheracademy.com/advent-2018/fmt/

package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

func main() {
	PrintV()
}

type Point struct {
	X int
	Y int
}

func PrintV() {
	var e interface{} = 2.718
	// %v print value of variable, %T print type of variable
	fmt.Printf("e=%v (%T)\n", e, e)

	//width  %10d指定输出的宽度，也可以通过*来格式化指定宽度
	fmt.Printf("%10d\n", 353)
	fmt.Printf("%*d\n", 10, 353)

	//align
	nums := []int{12, 237, 3838, 3}
	size := alignSize(nums)
	for i, n := range nums {
		fmt.Printf("%02d %*d\n", i, size, n)
	}

	//如果在格式字符串内多次引用变量，可以使用%[n]通过位置引用，其中n是参数的索引（基于1）
	fmt.Printf("The price of %[1]s was $%[2]d. $%[2]d! imagine that.\n", "carrot", 23)

	// %v print a Go value. + prefix to print field names in a struct and with # prefix to print field names and types
	p := &Point{1, 2}
	fmt.Printf("%v %+v %#v \n", p, p, p)

	//要控制对象的打印，需要实现fmt.Formatter接口和可选的fmt.Stringer接口
	ai := &AuthInfo{
		Login:  "holos",
		ACL:    ReadACL | WriteACL,
		APIKey: "duck season",
	}
	fmt.Println(ai.String())
	fmt.Printf("ai %%s: %s\n", ai)
	fmt.Printf("ai %%q: %q\n", ai)
	fmt.Printf("ai %%v: %v\n", ai)
	fmt.Printf("ai %%+v: %+v\n", ai)
	fmt.Printf("ai %%#v: %#v\n", ai)
}

type AuthInfo struct {
	Login  string //Login User
	ACL    uint   // ACL bitmask
	APIKey string // API key
}

const (
	keyMask  = "******"
	ReadACL  = 0x0000001
	WriteACL = 0x000001
)

//String implements Stringer interface
func (ai *AuthInfo) String() string {
	key := ai.APIKey
	if key != "" {
		key = keyMask
	}
	return fmt.Sprintf("Login:%s, ACL:%08b, APIKey: %s", ai.Login, ai.ACL, key)
}

var authInfoFields []string

func init() {
	typ := reflect.TypeOf(AuthInfo{})
	authInfoFields = make([]string, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		authInfoFields[i] = typ.Field(i).Name
	}
	sort.Strings(authInfoFields)
}

func (ai *AuthInfo) Format(state fmt.State, verb rune) {
	switch verb {
	case 's', 'q':
		val := ai.String()
		if verb == 'q' {
			val = fmt.Sprintf("%q", val)
		}
		fmt.Fprint(state, val)
	case 'v':
		if state.Flag('#') {
			fmt.Fprintf(state, "%T", ai)
		}
		fmt.Fprint(state, "{")
		val := reflect.ValueOf(*ai)
		for i, name := range authInfoFields {
			if state.Flag('#') || state.Flag('+') {
				fmt.Fprintf(state, "%s:", name)
			}
			fld := val.FieldByName(name)
			if name == "APIKey" && fld.Len() > 0 {
				fmt.Fprint(state, keyMask)
			} else {
				fmt.Fprint(state, fld)
			}
			if i < len(authInfoFields)-1 {
				fmt.Fprint(state, " ")
			}
		}
		fmt.Fprint(state, "}")
	}
}

func alignSize(nums []int) int {
	size := 0
	for _, n := range nums {
		if s := int(math.Log10(float64(n))) + 1; s > size {
			size = s
		}
	}
	return size
}
