package util

import "time"

// *** sign 1 bit, timestamp 42 bits, shard_id 9 bits, seq_id 12 bits ***
// result := (now_millis - our_epoch) << 21;	-- Shift left by 21 bits to make room for shard_id + sequence_id
// result := result | (%d << 12);				-- Shift left by 12 bits to make room for sequence_id
// result := result | (seq_id); 				-- Keep sequence_id in the last 12 bits

const ourEpoch int64 = 1314220021721

// Decode ID function
func DecodeID(id int64) (time.Time, int64, int64) {
	timestamp := (id >> 21) + ourEpoch
	shardID := (id >> 12) & 511
	sequenceID := id & 4095

	// Convert timestamp to human-readable format
	t := time.UnixMilli(timestamp)

	return t, shardID, sequenceID
}

func ExtractShardID(id int64) int {
	// note: (1<<9)-1 = 511, like above
	return int((id>>12)&((1<<9)-1) - 1)
}
