package main

import "fmt"

// func canMakeSubsequence(str1 string, str2 string) bool {
// 	i, n := 0, len(str2)
// 	for _, c := range str1 {
// 		d := byte('a')

// 		if (c != 'z') {
// 			d = byte(c)+1
// 		}

// 		if i < n && (str2[i] == byte(c) || str2[i] == d) {
// 			i++
// 		}
// 	}
// 	return i==n
// }

func main() {
    m := "vansh"
    n:= len(m)
    for i:=0; i<n; i++{
        for j:=i+1; j<=n; j++ {
            fmt.Println(m[i:j])
        }
    }

}
