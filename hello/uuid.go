package hello

import "github.com/google/uuid"

func myuuid() {
	for i := 0; i < 5; i++ {
		println(uuid.New().String())
	}
}
