package main

import (
	"ClickHouse-Connector/configuration"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func main() {
	conn, err := connect()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	rows, err := conn.Query(ctx, "select * from test;")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var iccid string
		if err := rows.Scan(
			&iccid,
		); err != nil {
			log.Fatal(err)
		}
		log.Printf("iccid: %s", iccid)
	}

}

func connect() (driver.Conn, error) {

	cfg := configuration.Config{}
	var (
		ctx = context.Background()

		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{cfg.AddrClick + ":" + cfg.PortClick},
			Auth: clickhouse.Auth{
				Database: cfg.Database,
				Username: cfg.Username,
				Password: cfg.Password,
			},
			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "an-example-go-client", Version: "0.1"},
				},
			},

			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		var exception *clickhouse.Exception
		if errors.As(err, &exception) {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil
}
