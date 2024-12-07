package helper

import "log"

func PanicIfError(err error) {
	if err != nil {
		log.Printf("Error occurred: %v", err) // Tambahkan log
		panic(err)

	}

}
