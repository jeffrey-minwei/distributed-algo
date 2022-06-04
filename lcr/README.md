# Experiments on LCR algorithm

## Implements with N node
for example, N = 5
```
$ go run lcrNodeN.go 
5 
name: 4 , uid: 595 , received: 905 , status: unknown
name: 3 , uid: 905 , received: 922 , status: unknown
name: 0 , uid: 939 , received: 595 , status: unknown
name: 2 , uid: 922 , received: 315 , status: unknown
name: 1 , uid: 315 , received: 939 , status: unknown
round 1 Done

name: 4 , uid: 595 , received: 922 , status: unknown
name: 1 , uid: 315 , received: 939 , status: unknown
name: 0 , uid: 939 , received: 905 , status: unknown
name: 2 , uid: 922 , received: 939 , status: unknown
name: 3 , uid: 905 , received: 922 , status: unknown
round 2 Done

name: 4 , uid: 595 , received: 922 , status: unknown
name: 1 , uid: 315 , received: 939 , status: unknown
name: 0 , uid: 939 , received: 922 , status: unknown
name: 3 , uid: 905 , received: 939 , status: unknown
name: 2 , uid: 922 , received: 939 , status: unknown
round 3 Done

name: 0 , uid: 939 , received: 922 , status: unknown
name: 2 , uid: 922 , received: 939 , status: unknown
name: 3 , uid: 905 , received: 939 , status: unknown
name: 4 , uid: 595 , received: 939 , status: unknown
name: 1 , uid: 315 , received: 939 , status: unknown
round 4 Done

name: 4 , uid: 595 , received: 939 , status: unknown
name: 0 , uid: 939 , received: 939 , status: leader
name: 1 , uid: 315 , received: 939 , status: unknown
name: 2 , uid: 922 , received: 939 , status: unknown
name: 3 , uid: 905 , received: 939 , status: unknown
round 5 Done
```
