package main

import (
	"context"
	"fmt"
	"study/feature_postgres/simple_connection"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err = conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("succeed!")
}
