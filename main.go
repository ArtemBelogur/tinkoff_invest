package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vodolaz095/go-investAPI/investapi"
	"log"
	"time"
)

const token = "t.k36HhmnrEuXVzvaEHzxu27_1lpxsQAbeaE3EqhTcm3VAb343GGgdxVqIVHceZDJWpawJhXmXYMc12tDx72a5ew"

func main() {
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/tinkoff_invest")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	client, err := investapi.New(token)
	if err != nil {
		log.Fatalf("%s : при соединении с investAPI", err)
	}

	fmt.Println("Подключено к MySQL")

	for true {
		var values [5]float64

		res, err := client.MarketDataServiceClient.GetLastPrices(context.Background(),
			&investapi.GetLastPricesRequest{Figi: []string{"BBG00RRT3TX4", "BBG00RKDZWG3", "BBG002PD3452", "BBG003TTSBB1", "BBG00425VG07"}},
		)
		if err != nil {
			log.Fatalf("%s : при получении котировок инструмента ОФЗ 25084", err)
		}

		for i, price := range res.GetLastPrices() {
			values[i] = price.GetPrice().ToFloat64()
		}

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `obligations` (`time`, `ofz_24021`, `ofz_25084`, `ofz_26207`, `ofz_26211`, `ofz_26212`)"+
			" VALUES (SYSDATE(), '%.4f', '%.4f', '%.4f', '%.4f', '%.4f')", values[0], values[1], values[2], values[3], values[4]))

		if err != nil {
			panic(err)
		}

		defer insert.Close()

		time.Sleep(10 * time.Minute)
	}
	err = client.Connection.Close()
	if err != nil {
		log.Fatalf("%s : при закрытии соединения", err)
	}
}
