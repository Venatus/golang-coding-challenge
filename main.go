package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

var DB driver.Conn

func main() {

	if err := connectClickhouse(); err != nil {
		log.Fatalf("cannot connect to clickhouse: %s", err.Error())
	}

}

func connectClickhouse() (err error) {

	ctx := context.Background()
	DB, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, "tcp", addr)
		},
		Debug: false,
		Settings: clickhouse.Settings{
			"max_execution_time": 3600,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:          time.Second * 30,
		MaxOpenConns:         5,
		MaxIdleConns:         5,
		ConnMaxLifetime:      time.Duration(60) * time.Minute,
		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
		ClientInfo: clickhouse.ClientInfo{
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "clickhouse-migrate", Version: "0.1"},
			},
		},
	})

	if err != nil {
		return err
	}

	if err := DB.Ping(ctx); err != nil {
		return err
	}

	log.Println("Connected to Clickhouse.")

	return nil
}
