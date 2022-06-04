package main

import (
    "fmt"
    "sync"
)

type Node struct {
    name string
    from chan int
    to chan int
    send int
    uid int
    status string
}

func send(n Node) {
    n.to <- n.send
}

func receive(n* Node, wg *sync.WaitGroup) {
    defer wg.Done()

    v := <-(n.from)
    if(v > n.uid) {
        n.send = v
    }else if(v == n.uid) {
        n.status = "leader"
    }
    fmt.Println("name:", n.name, ", uid:", n.uid, ", received:", v, ", status:", n.status)
}

/**
 * go run lcr.go
 * https://hackmd.io/@butastur/lcr-algorithm
 */
func main() {

    n1 := Node{name: "n1", uid: 765, send: 765, to: make(chan int), status: "unknown"}
    n2 := Node{name: "n2", uid: 12, send: 12, to: make(chan int), status: "unknown"}
    n3 := Node{name: "n3", uid: 65, send: 65, to: make(chan int), status: "unknown"}
    n4 := Node{name: "n4", uid: 123, send: 123, to: make(chan int), status: "unknown"}

    n2.from = n1.to
    n3.from = n2.to
    n4.from = n3.to
    n1.from = n4.to

    wg := new(sync.WaitGroup)

    round := 0
    for true {
        wg.Add(4)
        go send(n1) // 1->2
        go send(n2) // 2->3
        go send(n3) // 3->4
        go send(n4) // 4->1

        go receive(&n1, wg)
        go receive(&n2, wg)
        go receive(&n3, wg)
        go receive(&n4, wg)
        wg.Wait()

        // next round
        round++
        fmt.Println("round", round, "Done")
        fmt.Println()

        if(n1.status == "leader" || n2.status == "leader" || n3.status == "leader" || n4.status == "leader") {
            break;
        }
    }
}
