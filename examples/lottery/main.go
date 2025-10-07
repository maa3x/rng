package main

import (
	"fmt"

	"github.com/maa3x/rng"
)

func main() {
	l := rng.NewLottery[string]()
	draw(l)

	l = rng.NewLottery[string]()
	l.Append("Alice").AppendWithWeight(2.5, "Charlie").AppendWithWeight(0.25, "Dave").AppendWithWeight(0.1, "Eve").AppendWithWeight(0.02, "Grace")
	draw(l)

	l = rng.NewLottery[string]()
	l.Append("Alice").AppendWithWeight(2.0, "Charlie").AppendWithWeight(1.5, "Frank")
	draw(l)

	l = rng.NewLottery[string]()
	l.AppendWithWeight(2.0, "Alice").AppendWithWeight(0.003, "Charlie").AppendWithWeight(1e-5, "Frank")
	draw(l)

	l = rng.NewLottery[string]()
	l.AppendWithWeight(0.25, "Alice").AppendWithWeight(2.0, "Charlie").AppendWithWeight(0, "Dave")
	draw(l)

	l = rng.NewLottery[string]()
	l.Append("Alice").AppendWithWeight(0.1, "Charlie").AppendWithWeight(0, "Dave")
	draw(l)

	l = rng.NewLottery("Alice", "Bob", "Charlie")
	draw(l)

	l = rng.NewLottery("Dave").AppendWithWeight(-1, "Alice", "Bob", "Charlie")
	draw(l)

	l = rng.NewLottery[string]()
	l.AppendWithWeight(-1, "Alice", "Bob", "Charlie")
	draw(l)

	l = rng.NewLottery[string]()
	l.AppendWithWeight(0.1, "Alice", "Bob", "Charlie")
	draw(l)
}

func draw(l *rng.Lottery[string]) {
	for range 1000000 {
		l.Draw()
	}
	fmt.Println(l.Items())
}
