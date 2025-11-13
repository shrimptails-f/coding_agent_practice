package main

import (
	"business/internal/library/mysql"
	"business/tools/seeder/seeders"
	"errors"
	"fmt"
	"os"

	"gorm.io/gorm"
)

func main() {
	if err := checkArgs(); err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	env := os.Args[1]
	conn, cleanup, err := openConnection(env)
	if cleanup != nil {
		defer cleanup()
	}
	if err != nil {
		panic(err)
	}

	if conn == nil || conn.DB == nil {
		panic("データベース接続が初期化されていません。")
	}

	tx, done := mysql.Transactional(conn.DB)
	defer done()

	if err := seed(tx); err != nil {
		tx.Error = err
		fmt.Printf("データ投入中にエラーが発生しました: %v\n", err)
		return
	}

	fmt.Println("正常に終了しました。")
}

func checkArgs() error {
	if len(os.Args) != 2 {
		return errors.New("期待している引数は1つです。引数を確認してください。")
	}

	if os.Args[1] != "dev" && os.Args[1] != "test" {
		return errors.New("第一引数が期待している語群は以下の通りです。\n1:dev\n2:test")
	}

	return nil
}

func openConnection(env string) (*mysql.MySQL, func() error, error) {
	switch env {
	case "dev":
		conn, err := mysql.New()
		return conn, nil, err
	case "test":
		conn, err := mysql.NewTest()
		return conn, nil, err
	default:
		return nil, nil, fmt.Errorf("unsupported env: %s", env)
	}
}

func seed(tx *gorm.DB) error {
	return seeders.CreateSamples(tx)
}
