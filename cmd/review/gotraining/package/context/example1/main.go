//这个例子讲了2个问题
//parent的cancel会影响child context  第一个
//only for request-scoped data   第二个
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type testcontext struct{}

func main() {
	test1234 := testcontext{}

	ctx := context.WithValue(context.Background(), test1234, "12345")

	ctx2, cancel := context.WithTimeout(ctx, 1*time.Nanosecond)

	var wg sync.WaitGroup
	cancel()

	wg.Add(1)

	go func() {

		test123 := testcontext{}

		time.Sleep(1 * time.Second)

		ctx3 := context.WithValue(ctx2, test123, "123")

		select {
		case <-ctx3.Done():
			{
				fmt.Println("canceled")

				fmt.Println(ctx3.Value(test1234).(string))

				fmt.Println(ctx3.Value(test123).(string))
			}
		}

		wg.Done()

	}()

	fmt.Println(ctx2.Value(test1234).(string))

	wg.Wait()

}
