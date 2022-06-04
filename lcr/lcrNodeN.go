package main

import (
    "fmt"
    "sync"
    "bufio"
    "os"
    "strconv"
    "time"
    "math/rand"
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
 * go run lcrNodeN.go
 * https://hackmd.io/@butastur/lcr-algorithm
 */
func main() {

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    sizeStr := scanner.Text()
    size, err := strconv.Atoi(sizeStr)
    if err != nil {
        fmt.Printf("%s", err.Error())
    }
    nodes := make([]Node, size)

    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        name := strconv.Itoa(i)
        // TODO: this uid is not absolutely unique, it have to be implemented
        uid := rand.Intn(400) + rand.Intn(999) + rand.Intn(200)
        nodes[i] = Node{name: name, uid: uid, send: uid, to: make(chan int), status: "unknown"}
    }

    for i := 1; i < size; i++ {
        nodes[i].from = nodes[i-1].to
    }
    nodes[0].from = nodes[size-1].to

    wg := new(sync.WaitGroup)

    leader_elect_done := false
    round := 0
    for ! leader_elect_done {
        wg.Add(size)

        for i := 0; i < size; i++ {
            go send(nodes[i])
        }

        for i := 0; i < size; i++ {
            go receive(&nodes[i], wg)
        }
        wg.Wait()

        // next round
        round++
        fmt.Println("round", round, "Done")
        fmt.Println()

        for i := 0; i < size; i++ {
            if(nodes[i].status == "leader") {
                leader_elect_done = true
                break;
            }
        }
    }
}
