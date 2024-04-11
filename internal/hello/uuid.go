package hello

import "github.com/google/uuid"

func Myuuid() {
	for i := 0; i < 5; i++ {
		println(uuid.New().String())
	}
}
