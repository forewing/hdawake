package hdawake

import (
	"fmt"
	"os"
	"time"
)

func Once(name string) error {
	data := []byte(time.Now().String())
	return os.WriteFile(name, data, 0777)
}

func Keep(name string, period time.Duration) {
	fmt.Printf("start writting to `%s` every %s\n", name, period)
	for {
		if err := Once(name); err == nil {
			fmt.Printf("%s\tsuccess\n", time.Now())
		} else {
			fmt.Printf("%s\terror: %s\n", time.Now(), err)
		}
		time.Sleep(period)
	}
}
