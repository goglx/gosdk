package s3

import "fmt"

const s3URL = "https://%s.s3.%s.amazonaws.com/%s"

func debug() string {
	return fmt.Sprintf(s3URL, "1", "2", "3")
}
