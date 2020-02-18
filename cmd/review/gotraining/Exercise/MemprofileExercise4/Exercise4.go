package MemprofileExercise4

import (
	"math/rand"
	"runtime"
	"sync"
)

func Exercise4(){
	ch := make(chan int,1000)

	var wg sync.WaitGroup
	poolsize := runtime.NumCPU()
	wg.Add(poolsize)

	shutdown := make(chan struct{})


	for i:=0;i<poolsize;i++{
				go func(id int) {
					for {
						num := rand.Intn(1000)

						select{
						case ch <- num:
							//fmt.Printf("Worker %d sent %d\n", id, num)
						case  <- shutdown:
							//fmt.Printf("Worker %d shutdown\n", id)
							wg.Done()
							return
						}
					}
				}(i)
			}
	var nums []int
	for n := range ch {
		if n%2 == 0{
			continue
		}
		nums = append(nums,n)

		if len(nums) == 100{
			break
		}
	}
	close(shutdown)
	wg.Wait()
	//fmt.Println(nums)
}

func Exercise5(){
	ch := make(chan int)

	var wg sync.WaitGroup
	poolsize := runtime.NumCPU()
	wg.Add(poolsize)

	shutdown := make(chan struct{})


	for i:=0;i<poolsize;i++{
		go func(id int) {
			for {
				num := rand.Intn(1000)

				select{
				case ch <- num:
					//fmt.Printf("Worker %d sent %d\n", id, num)
				case  <- shutdown:
					//fmt.Printf("Worker %d shutdown\n", id)
					wg.Done()
					return
				}
			}
		}(i)
	}
	var nums []int
	for n := range ch {
		if n%2 == 0{
			continue
		}
		nums = append(nums,n)

		if len(nums) == 100{
			break
		}
	}
	close(shutdown)
	wg.Wait()
	//fmt.Println(nums)
}