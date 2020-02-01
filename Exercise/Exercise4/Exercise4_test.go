package Exercise4

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