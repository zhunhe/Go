package main

/*
#include <stdio.h>
void print(char *str) {
	printf("%s\n", str);
}
*/
import "C"

func main() {
	yourName := "print string using c language"
	C.print(C.CString(yourName))
}
