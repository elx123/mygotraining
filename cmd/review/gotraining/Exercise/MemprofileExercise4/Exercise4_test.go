package MemprofileExercise4
//benchmark memory 通过相同的时间对比操作的次数,与视频互相呼应
//   go test -gcflags "-m -m" -run none -bench . -benchtime 3s -benchmem -memprofile m.out
//	通过http查看
//	go tool pprof -http :3000 c.out
//	得出的结果是Exercise4使用buffer channel分配了更多的内存指的是alloc_space
import (
	"testing"
)

func BenchmarkExercise4(b *testing.B) {
	for i:=0;i<b.N;i++{
		Exercise4()
	}
}

func BenchmarkExercise5(b *testing.B) {
	for i:=0;i<b.N;i++{
		Exercise5()
	}
}