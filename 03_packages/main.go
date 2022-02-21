package main

import (
	"fmt"
	"math"

	"github.com/avin-dsilva-bazaarvoice/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(22.7))
	fmt.Println(math.Ceil(22.7))
	fmt.Println(math.Sqrt(22.7))
	fmt.Println(strutil.Reverse("hello"))
}
