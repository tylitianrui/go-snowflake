package snowflake

import (
	"fmt"
	"testing"
)

func BenchmarkSnowFlake_NextID(b *testing.B) {
	var (
		sf  *SnowFlake
		err error
		id  uint64
	)
	if sf, err = NewSnowFlake(12, 5); err != nil {
		fmt.Println(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if id, err = sf.NextID(); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(id)

	}

}
