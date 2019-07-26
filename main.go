package main

func sum(chA, chB <-chan int) <-chan int {
	res := make(chan int)
	queueA := make(queue)
	queueB := make(queue)

	go func () {
		for {
			select {
			case valA, ok := <-chA:
				if !ok {
					close(res)
					return
				}
				queueA.push(valA)
			case valB, ok := <-chB:
				if !ok {
					close(res)
					return
				}
				queueB.push(valB)
			}
			if !queueA.empty() && !queueB.empty() {
				res <- queueA.pop() + queueB.pop()
			}
		}
	}()

	return res
}

func main() {
	
}
