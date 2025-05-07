// bobの練習
// 2025 (c) Tadahiro Ishisaka
package main

import (
	"context"
	"fmt"
	"github.com/ishisaka/golang-bob-plactice/models"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/mysql"
	"github.com/stephenafamo/bob/dialect/mysql/sm"
	"github.com/stephenafamo/scan"
	"reflect"
)

// main関数は、MySQLデータベースに接続し、worldデータベースから都市情報を取得して表示します。
// 具体的には以下の処理を行います：
// - MySQLデータベースへの接続
// - ID=1の都市情報の取得と表示
// - ID>1の全都市情報の取得と一覧表示
func main() {
	// DBへの接続
	db, err := bob.Open("mysql", "root:root@tcp(localhost:3306)/world")
	if err != nil {
		panic(err)
	}

	// ビルダーによって作られたモデルを使う。
	// ID==1の都市を取得
	city, err := models.FindCity(context.Background(), db, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(city)

	//IDが1より大きい都市を全て取得
	cities, err := models.Cities.Query(
		models.SelectWhere.Cities.ID.GT(1),
	).All(context.Background(), db)
	if err != nil {
		panic(err)
	}
	for i, city := range cities {
		fmt.Printf("%d: %s\n", i, city)
	}

	// 生クエリ、SELECTの例
	q := mysql.Select(
		sm.Columns("ID", "Name", "CountryCode", "District", "Population"),
		sm.From("city"),
		sm.Where(mysql.Quote("Name").EQ(mysql.Arg("York"))),
	)
	result, err := bob.All(context.Background(), db, q, scan.StructMapper[models.City]())
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(result))
	for _, city := range result {
		fmt.Println(city)
	}
}
