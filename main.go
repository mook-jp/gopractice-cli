<<<<<<< HEAD
package main

import "fmt"

const (
	first1 = 200 + iota
	second1
)
const (
	first2 = iota
	second2
)

func main() {
	fmt.Println("vim-go")

	const HTTPStatusOK = 200

	const (
		StatusOk              = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)

	println("go begin")
	fmt.Printf("first1  = %d\n", first1)
	fmt.Printf("second1 = %d\n", second1)
	println("go end")
=======
/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import "goPractice/cmd"

func main() {
	cmd.Execute()
>>>>>>> a66fdd3 (cobra-cli init)
}
