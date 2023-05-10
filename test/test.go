package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(RobinFingerprint("KIM"));
}

func RobinFingerprint(username string) int{
	username = strings.ToLower(username);
	var sum float64 = 0;
	for i := 0; i < len(username); i++ {
		sum += float64((float64(username[i]) * math.Pow(16, float64(i))));
	}
	return int(math.Mod(sum, 64)) ;
 }