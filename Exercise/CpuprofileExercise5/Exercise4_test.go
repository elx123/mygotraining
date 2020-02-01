package CpuprofileExercise5
//benchmark memory 通过相同的时间对比操作的次数,与视频互相呼应
//  go test -run none -bench . -benchtime 3s -cpuprofile c.out
//	通过http查看
//	go tool pprof -http :3000 c.out
//  可以通过top和source查看代码,从而优化代码
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