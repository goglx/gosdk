package s3

import "fmt"

var s3URL = "https://%s.s3.%s.amazonaws.com/%s"

func Debug() {
	fmt.Println("s3 debug", s3URL)
}
