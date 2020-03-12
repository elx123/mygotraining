package listMy_test

import (
	listMy "mygotraining/cmd/review/gotraining/algorithms/data/listMy"
	"testing"
)

func TestListAdd(t *testing.T) {
	t.Helper()
	t.Log("we will test our list package")
	{
		t.Log("we will test the add function")
		{
			ls := listMy.List{}
			ls.Add("test1")
			n := ls.Add("test2")
			t.Log("we get the result")
			//t.Logf("the list first address %p,the list node address %p",ls.)
			t.Logf("%v %p", ls, n)
		}
	}
}

func TestListAddFront(t *testing.T) {
	t.Log("we will test our list package")
	{
		t.Log("we will test the AddFront function")
		{
			ls := listMy.List{}
			ls.AddFront("test1")
			n := ls.AddFront("test2")
			t.Log("we get the result")
			//t.Logf("the list first address %p,the list node address %p",ls.)
			t.Logf("%v %p", ls, n)
		}
	}
}

func TestListFind(t *testing.T) {
	t.Log("we will test our list package")
	{
		t.Log("we will test the Find function")
		{
			ls := listMy.List{}
			n := ls.AddFront("test1")
			ls.AddFront("test2")
			matchn, err := ls.Find("test1")
			if err != nil {
				t.Fatal(err)
			}
			t.Log("we get the result")
			//t.Logf("the list first address %p,the list node address %p",ls.)
			t.Logf("origin value %v,origin address %p,matched value %v,matched address %p", n, n, matchn, matchn)
		}
	}
}

func TestListFindResver(t *testing.T) {
	t.Log("we will test our list package")
	{
		t.Log("we will test the FindResver function")
		{
			ls := listMy.List{}
			n := ls.AddFront("test1")
			ls.AddFront("test2")
			matchn, err := ls.FindResver("test1")
			if err != nil {
				t.Fatal(err)
			}
			t.Log("we get the result")
			//t.Logf("the list first address %p,the list node address %p",ls.)
			t.Logf("origin value %v,origin address %p,matched value %v,matched address %p", n, n, matchn, matchn)
		}
	}
}

func TestRemove(t *testing.T) {
	t.Log("we will test our list package")
	{
		t.Log("we will test the Remove function")
		{
			ls := listMy.List{}
			ls.AddFront("test1")
			ls.AddFront("test2")
			_, err := ls.Remove("test1")
			if err != nil {
				t.Fatal(err)
			}
			_, err = ls.Find("test1")
			if err != nil {
				t.Log("remove success")
				return
			}
			t.Fail()
		}
	}
}
