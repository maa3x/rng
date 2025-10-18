package main

import (
	"fmt"

	"github.com/maa3x/rng"
)

func main() {
	l := rng.NewLottery[string]()
	draw(l)

	l = rng.NewLottery[string]()
	l.Append("Alice").AppendWeight(2.5, "Charlie").AppendWeight(0.25, "Dave").AppendWeight(0.1, "Eve").AppendWeight(0.02, "Grace")
	draw(l)

	l = rng.NewLottery[string]()
	l.Append("Alice").AppendWeight(2.0, "Charlie").AppendWeight(1.5, "Frank")
	draw(l)

	l = rng.NewLottery[string]()
	l.AppendWeight(2.0, "Alice").AppendWeight(0.003, "Charlie").AppendWeight(1e-5, "Frank")
	draw(l)

	l = rng.NewLottery[string]()
	l.AppendWeight(0.25, "Alice").AppendWeight(2.0, "Charlie").AppendWeight(0, "Dave")
	draw(l)

	l = rng.NewLottery[string]()
	l.Append("Alice").AppendWeight(0.1, "Charlie").AppendWeight(0, "Dave")
	draw(l)

	l = rng.NewLottery("Alice", "Bob", "Charlie")
	draw(l)

	l = rng.NewLottery("Dave").AppendWeight(-1, "Alice", "Bob", "Charlie")
	draw(l)

	l = rng.NewLottery[string]()
	l.AppendWeight(-1, "Alice", "Bob", "Charlie")
	draw(l)

	l = rng.NewLottery[string]()
	l.AppendWeight(0.1, "Alice", "Bob", "Charlie")
	draw(l)
}

func draw(l *rng.Lottery[string]) {
	for range 1000000 {
		l.Draw()
	}
	fmt.Println(l.Items())
}
