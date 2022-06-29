package main

import (  
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
)

type bill struct {
	id uint
	timestamp time.Time
	in_ex string
	commodity string
	commodity_type string
	ammount decimal.Decimal
	pay_method string
	counterparty string
	tran_account string
	order_number string
	mer_number string
	source string
	remark string
}

func main() {
   // 准备数据库连接
	db, err := sql.Open("mysql", "root:Ltf*ms93@tcp(127.0.0.1:3306)/bill?parseTime=true")  
	if err != nil {
		log.Fatal(err)
	}
	
	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 插入数据
	// stmt, err := db.Prepare("INSERT INTO bill(timestamp, in_ex, commodity, commodity_type, ammount, pay_method, counterparty, tran_account, order_number, mer_number, source, remark)VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ammount, err := decimal.NewFromString("123.45")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// res, err := stmt.Exec(time.Now(), "收入", "水杯", "日用百货", ammount, "中国邮政银行储蓄卡", "xxx", "liyg@163.com", "2022062777001178281442785564", "T200P2725499089050050533", "支付宝", "水杯")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// lastId, err := res.LastInsertId()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// rowCnt, err := res.RowsAffected()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	// 查询数据
	stmt, err := db.Prepare("select * from bill")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	t := &bill{}
	for rows.Next() {
		err := rows.Scan(&t.id, &t.timestamp, &t.in_ex, &t.commodity, &t.commodity_type, &t.ammount, &t.pay_method, &t.counterparty, &t.tran_account, &t.order_number, &t.mer_number, &t.source, &t.remark)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(t)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}