## Week02 pull request is about the unit test cases for the composite types like arrays, slices, and maps. For testing unit test cases, imported the testing package and used the Error,Errorf, fatal and Fatalf functions from the testing package.

<br />

## Array's first test case is about testing the contents of the two array types.
	an exp array variable is initialized with four strings and an act array variable is initialized with empty. Range loop is used for iterating through the exp array variable and copy exp values into the act array variable. Range loop is used for iterating through the act array variable and checking the act array variable content with the exp array variable content. If both array contents are not equal, the test case will fail and will output the error message using the Errorf function from the testing package. If both array contents are equal, the test case will pass.	

## Array's second test case is about testing the contents of the two array types.
	an exp array variable is initialized with four strings and an act array variable is initialized with empty. Range loop is used for iterating through the exp array variable and copy exp values into the act array variable. DeepEqual function from reflect package is used for checking the contents of act array variable and exp array variable. If both array contents are not equal, the test case will fail and will output the error message using the Error function from the testing package. If both array contents are equal, the test case will pass.
	
## Array's third test case is about testing the length of the two array types.
    an exp array variable is initialized with four strings and an act array variable is initialized with empty. Range loop is used for iterating through the exp array variable and copy exp values into the act array variable. len function from the builtin package is used for checking the length of the act array variable and exp array variable. If both array lengths are not equal, the test case will fail and will output the error message using the Errorf function from the testing package. If both array contents are equal, the test case will pass.

## Slice's first test case is about testing the contents of the two slice types.
    exp slice variable is initialized with four strings and act slice variable is initialized using make function from builtin package with length=0 and capacity=length of exp. Range loop is used for iterating through the exp slice variable and copy exp values into the act slice variable. Range loop is used for iterating through the act slice variable and checking the act slice variable content with the exp slice variable content. If both slice contents are not equal, the test case will fail and will output the error message using the Errorf function from the testing package. If both slice contents are equal, the test case will pass.

## Slice's second test case is about testing the contents of the two slice types.
    exp slice variable is initialized with four strings and act slice variable is initialized with empty. Range loop is used for iterating through the exp slice variable and copy exp values into the act slice variable. DeepEqual function from reflect package is used for iterating through the act slice variable and checking the act slice variable contents with the exp slice variable content. If both slice contents are not equal, the test case will fail and will output the error message using the Error function from the testing package. If both slice contents are equal, the test case will pass.	
	
## Slice's third test case is about testing the length of the two slice types.
    exp slice variable is initialized with four strings and act slice variable is initialized with empty. Range loop is used for iterating through the exp slice variable and copy exp values into the act slice variable. len function from the builtin package is used for checking the length of the act slice variable and exp slice variable. If both slice lengths are not equal, the test case will fail and will output the error message using the Errorf function from the testing package. If both slice contents are equal, the test case will pass.
	
## Slice's fourth test case is about testing the length and content of the two slice types.
    exp slice variable is initialized with four strings and act slice variable is initialized with empty. Range loop is used for iterating through the exp slice variable and copy exp values into the act slice variable. len function from the builtin package is used for checking the length of the act slice variable and exp slice variable. If both slice lengths are not equal, the test case will fail and will output the error message using the Errorf function from the testing package. Range loop is used for iterating through the act slice variable and checking the act slice variable content with the exp slice variable content. If both slice contents are not equal, the test case will fail and will output the error message using the Errorf function from the testing package. If both slice contents are equal, the test case will pass.
	
## Map first test case is about testing the contents of the two map types.
    exp map variable is initialized with four strings and act map variable is initialized using make function from builtin package with length=0 and capacity=length of exp. Range loop is used for iterating through the exp map variable and copy exp values into the act map variable. Range loop is used for iterating through the act map variable and checking the act map variable content with the exp map variable content. If both map contents are not equal, the test case will fail and will output the error message using the Errorf function from the testing package. If both map contents are equal, the test case will pass.

## Map second test case is about testing the contents of the two map types.
    exp map variable is initialized with four strings and act map variable is initialized with empty. Range loop is used for iterating through the exp map variable and copy exp values into the act map variable.
	DeepEqual function from reflect package is used for iterating through the act map variable and checking the act map variable content with the exp map variable content. If both map contents are not equal, the test case will fail and will output the error message using the Error function from the testing package. If both map contents are equal, the test case will pass.

## Map third test case is about testing the length of the two map types.
	exp map variable is initialized with four strings and act map variable is initialized with empty. Range loop is used for iterating through the exp map variable and copy exp values into the act map variable. len function from the builtin package is used for checking the length of the act map variable and exp map variable. If both map contents are not equal, the test case will fail and will output the error message using the Errorf function from the testing package. If both map contents are equal, the test case will pass.

## Map fourth test case is about testing the length and content of the two map types.
    exp map variable is initialized with four strings and act map variable is initialized with empty. len function from the builtin package is used for checking the length of the act map variable and exp map variable. If both map lengths are not equal, the test case will fail and will output the error message using the Errorf function from the testing package. Range loop is used for iterating through the act map variable and checking the act map variable content with the exp map variable content. If both map contents are not equal, the test case will fail and will output the error message using the Errorf function from the testing package. If both map contents are equal, the test case will pass.
	

# Testing
### Error, ErrorF, Fatal, FatalF functions from the testing package. With Error and Errorf test case is marked as failed and the test function continues running. With Fatal and Fatalf test case is marked as failed and the test function exits immediately. Any remaining test functions will execute after the current test function exits.

<br />

```sh
func Test_Error_1(t *testing.T) {
	a1 := [4]string{"John", "Paul", "George", "Ringo"}
	a2 := [4]string{"John", "Paul", "Peter", "David"}
	for i, v := range a1 {
		if v != a2[i] {
			t.Errorf("act[%d] : %s is not equal to exp[%d] : %s ", i, v, i, a2[i])
		}
	}
}
=== RUN   Test_Error_1
    assignment02_test.go:15: act[2] : George is not equal to exp[2] : Peter 
    assignment02_test.go:15: act[3] : Ringo is not equal to exp[3] : David 
--- FAIL: Test_Error_1 (0.00s)
```
```sh
func Test_Fatal_1(t *testing.T) {
	a1 := [4]string{"John", "Paul", "George", "Ringo"}
	a2 := [4]string{"John", "Paul", "Peter", "David"}
	for i, v := range a1 {
		if v != a2[i] {
			t.Fatalf("act[%d] : %s is not equal to exp[%d] : %s ", i, v, i, a2[i])
		}
	}
}
=== RUN   Test_Fatal_1
    assignment02_test.go:27: act[2] : George is not equal to exp[2] : Peter 
--- FAIL: Test_Fatal_1 (0.00s)
```

<br />

# Below are the interesting code snippets which surprised me and excited me also

<br />

```sh
i9 := []int{1, 2, 3, 4}
i10 := []int{1, 2, 3, 4}
fmt.Println("Contents of i5 and i6 are : ", i9 == i10) ---> Compile time error : invalid operation: i9 == i10 (slice can only be compared to nil)
```

```sh
i12 := []int{4,5,6}
append(i12, 7) ---> Compile Time Error : append(i12, 7) evaluated but not used
```

```sh
i13 := []int{4,5,6}
i14 := []int{7,8,9}
i13 = append(i13, i14)  ---> Compile Time Error : cannot use i14 (type []int) as type int in append
fmt.Println("i13 : ", i13)
```

```sh
i13 := []int{4, 5, 6}
i14 := []int{7, 8, 9}
i13 = append(i13, i14...)
fmt.Println("i13 : ", i13)
```

```sh
x,y := 1,2
k := x*y
i19 := make([]int, 5, k) ---> panic: runtime error: makeslice: cap out of range
fmt.Println("i19 : ", i19)
```

```sh
var m1 map[string]int ---> nil map literal
m1["hello"] = 999       ---> panic: assignment to entry in nil map
```

```sh
m3 := map[string]int{} ---> empty map literal
```

```sh
m4 := map[int]string{  ---> non-empty map literal
		1: "Jan",
		2: "Feb",
		3: "Mar",  ---> comma is required on the last line also		
	}
fmt.Println("capacity of m4 : ", cap(m4)) ---> compile time error : invalid argument m4 (type map[int]string) for cap
```

```sh
var m7 map[int]string
var m8 map[int]string
fmt.Println(m7 == m8) ---> compile time error : invalid operation: m7 == m8 (map can only be compared to nil)
```

```sh
m9 := map[int]string{}
m9[1] = "Jan"
m9[2] := "Feb" ---> Compile time error : non-name m9[2] on left side of :=
```

```sh
m15 := map[func()]int{} ---> Compile time error : invalid map key type func()
fmt.Println(m15)
```

```sh
m15 := map[[]int]int{} ---> Compile time error : invalid map key type []int
fmt.Println(m15)
```

```sh
m15 := map[map[int]string]int{} ---> Compile time error : invalid map key type map[int]string
fmt.Println(m15)
```

```sh
#Using Maps as sets using booleans
s1 := []int{1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 8, 9, 10}	
m13 := map[int]bool{}
for _, v := range s1 {
	m13[v] = true  ---> Occupies 1 byte
}	
if m13[5] {
	fmt.Println("5 is present in set")
}
```

```sh
#Using Maps as sets using structs
s2 := []int{1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 8, 9, 10}	
m14 := map[int]struct{}{}
for _, v := range s2 {
	m14[v] = struct{}{} ---> Occupies 0 byte
}
if _, ok := m14[5]; ok {
	fmt.Println("5 is present in set")
}
```

-	### Maps are NOT concurrent safe.

-	### The comma ok idiom is used in Go when we want to differentiate between reading a value and getting back the zero value.

-	### Use a map when the order of elements doesn't matter. Use a slice when the order of elements is important.

-	### The Go runtime is compiled into every Go binary. This is different from languages that use a virtual machine, which must be installed separately to allow programs written in those languages to function. Including the runtime in the binary makes it easier to distribute Go programs and avoids worries about compatibility issues between the runtime and the program.