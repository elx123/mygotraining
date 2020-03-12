package internal

import "testing"

func ABC(t *testing.T){
	t.Helper()
	t.Log("1234")
}

func CBA(t *testing.T){
	t.Log("1234")
}
