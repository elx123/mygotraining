//only for request-scoped data   第二个
package main

import (
	"context"
	"fmt"
	"sync"
)

type testcontext struct{}

func main() {
	test1234 := testcontext{}

	test12345 := testcontext{}

	ctx1 := context.WithValue(context.Background(), test1234, "1234")

	ctx2 := context.WithValue(ctx1, test12345, "12345")

	fmt.Println(ctx1.Value(test1234).(string))

	fmt.Println(ctx2.Value(test12345).(string))

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		defer wg.Done()

		test123456 := testcontext{}

		ctx3 := context.WithValue(ctx2, test123456, "123456")

		fmt.Println(ctx3.Value(test1234).(string))

		fmt.Println(ctx3.Value(test12345).(string))
	}()

	wg.Wait()

}
