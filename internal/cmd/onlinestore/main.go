package main

//go:generate sqlboiler --wipe psql

import (
	"github.com/Abdurrochman25/online-store/internal/cmd"
	_ "github.com/lib/pq"
)

func main() {
	cmd.Init()
}
