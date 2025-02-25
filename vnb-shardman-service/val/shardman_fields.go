package val

import (
	"fmt"
)

func ValidateShardID(shardID int32) error {
	if shardID <= 0 {
		return fmt.Errorf("Shard ID must be greater than 0, starting from 1")
	}
	return nil
}
