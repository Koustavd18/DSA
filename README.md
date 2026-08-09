# 🧮 DSA

Data Structures & Algorithms practice in Go — one package per problem, each
exploring a **brute-force** solution and one or more **optimal** follow-ups.

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)
![Status](https://img.shields.io/badge/status-active-brightgreen)

## 🚀 Running

Each package exposes its solution functions and a test fixture (`Test` / a
package-level var). `main.go` wires up whichever package is currently being
explored:

```bash
go run .
```

Swap the import in [`main.go`](main.go) to try a different package.

## 📚 Problems

| Package | Problem | Approach(es) | Time | Space |
|---|---|---|---|---|
| [`BadVersion`](BadVersion) | First Bad Version | Binary search | `O(log n)` | `O(1)` |
| [`Binary`](Binary) | Add Binary Strings | Digit-by-digit with carry | `O(max(n,m))` | `O(max(n,m))` |
| [`Boat`](Boat) | Boats to Save People | Sort + two pointers (greedy) | `O(n log n)` | `O(1)` |
| [`CountPrimes`](CountPrimes) | Count Primes | Brute force trial division | `O(n²)` | `O(1)` |
| [`FirstAndLast`](FirstAndLast) | First & Last Position in Sorted Array | Linear scan → binary search | `O(n)` → `O(log n)` | `O(1)` |
| [`hasDuplicates`](hasDuplicates) | Contains Duplicate | Nested loop → sort → hash set | `O(n²)` → `O(n log n)` → `O(n)` | `O(1)` → `O(1)` → `O(n)` |
| [`longSubStr`](longSubStr) | Longest Substring Without Repeating Characters | Brute force → sliding window | `O(n²)` → `O(n)` | `O(min(n,charset))` |
| [`majority`](majority) | Majority Element | Nested loop → hash map → Boyer–Moore voting | `O(n²)` → `O(n)` → `O(n)` | `O(1)` → `O(n)` → `O(1)` |
| [`MissingNumber`](MissingNumber) | Missing Number | Sort & scan → Gauss sum | `O(n log n)` → `O(n)` | `O(1)` |
| [`MoveZeros`](MoveZeros) | Move Zeroes | Extra array → in-place two pointer | `O(n)` | `O(n)` → `O(1)` |
| [`ProductExpectSelf`](ProductExpectSelf) | Product of Array Except Self | Prefix × suffix products | `O(n)` | `O(1)` extra |
| [`Robot`](Robot) | Robot Return to Origin | Track x/y displacement | `O(n)` | `O(1)` |
| [`SingleNumber`](SingleNumber) | Single Number | Hash map counts → XOR | `O(n)` | `O(n)` → `O(1)` |
| [`TwoSum`](TwoSum) | Two Sum | Nested loop → hash map | `O(n²)` → `O(n)` | `O(1)` → `O(n)` |
| [`validMountain`](validMountain) | Valid Mountain Array | Single ascend/descend scan | `O(n)` | `O(1)` |
| [`waterContainer`](waterContainer) | Container With Most Water | Brute force → two pointers | `O(n²)` → `O(n)` | `O(1)` |

## 🗂 Layout

```
DSA/
├── main.go            # entry point, wires up the package under exploration
├── <Problem>/
│   └── *.go            # BruteForce / Optimal (+ variants) implementations
└── go.mod
```

## 🛠 Linting

This repo is checked with [Trunk](https://trunk.io) (`gofmt`, `golangci-lint2`,
`markdownlint`, `osv-scanner`, `trufflehog`). Run `trunk check` before
committing.
