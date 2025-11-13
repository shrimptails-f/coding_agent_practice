package main

import (
	"business/internal/library/mysql"
	"business/tools/migrations/model"
	"errors"
	"fmt"
	"os"
)

// main は引数からテーブル作成を行います
// 引数:
// - arg1: 接続環境の指定。期待する語群:dev or test
// - arg2: テーブルの作成か、削除の指定 期待する語群:create or drop
func main() {
	if err := checkArgs(); err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	env := os.Args[1]
	action := os.Args[2]

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

	switch action {
	case "create":
		err = conn.DB.AutoMigrate(migrationTargets()...)
	case "drop":
		err = conn.DB.Migrator().DropTable(migrationTargets()...)
	}

	if err != nil {
		fmt.Printf("エラーが発生しました。:%v\n", err)
		return
	}

	fmt.Println("正常に終了しました。")
}

// checkArgs はコマンドライン引数を確認する。
func checkArgs() error {
	if len(os.Args) != 3 {
		return errors.New("期待している引数は2つです。引数を確認してください。")
	}

	if os.Args[1] != "dev" && os.Args[1] != "test" {
		return errors.New("第一引数が期待している語群は以下の通りです。\n1:dev\n2:test")
	}

	if os.Args[2] != "create" && os.Args[2] != "drop" {
		return errors.New("第二引数が期待している語群は以下の通りです。\n1:create\n2:drop")
	}

	return nil
}

func migrationTargets() []interface{} {
	return []interface{}{
		model.Sample{},
	}
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
