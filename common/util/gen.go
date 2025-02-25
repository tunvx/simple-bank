package util

import "fmt"

func GenShardID(id int64, volume int64) int32 {
	fmt.Println("ID = ", id)
	fmt.Println("volume = ", volume)
	updatedShardID := (id-1)/volume + 1
	return int32(updatedShardID)
}
