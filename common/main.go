package main

import (
	"fmt"

	"github.com/tunvx/simplebank/common/util"
)

// Giải mã và hiển thị thông tin ID
func printDecodedID(id int64) {
	timestamp, shardID, sequenceID := util.DecodeID(id)
	fmt.Printf("ID: %d\n", id)
	fmt.Printf("  - Timestamp: %s\n", timestamp.Format("2006-01-02 15:04:05.000"))
	fmt.Printf("  - Shard ID: %d\n", shardID)
	fmt.Printf("  - Sequence ID: %d\n", sequenceID)
}

// Hiển thị shard ID từ ID
func printShardID(id int64) {
	shardID := util.ExtractShardID(id)
	fmt.Printf("ID: %d -> Shard ID: %d\n", id, shardID)
}

func main() {
	// Danh sách ID mẫu
	ids := []int64{
		892857634755973170, // shard = 1
		892857634823086081, // shard = 2
	}

	// Giải mã và in thông tin cho từng ID
	for _, id := range ids {
		printDecodedID(id)
		printShardID(id)
		fmt.Println() // Xuống dòng để dễ đọc
	}
}
