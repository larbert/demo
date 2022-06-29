package main

import (  
	"database/sql"
	"log"
	"time"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
    "github.com/xuri/excelize/v2"
)

type bill struct {
	Id uint
	Timestamp time.Time
	In_ex string
	Commodity string
	Commodity_type string
	Ammount decimal.Decimal
	Pay_method string
	Counterparty string
	Tran_account string
	Order_number string
	Mer_number string
	Source string
	Remark string
}

func main() {
	readWeiChatBill()
	readAlipayBill()
}

func readWeiChatBill() {
	// 准备数据库连接
	db, err := sql.Open("mysql", "root:Ltf*ms93@tcp(127.0.0.1:3306)/bill?parseTime=true")  
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO bill(timestamp, in_ex, commodity, commodity_type, ammount, pay_method, counterparty, tran_account, order_number, mer_number, source, remark)VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}


	// 打开excel文件
	f, err := excelize.OpenFile("excel1.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}
	rows = rows[17:]

	// 解析成bill对象
	bills := make([]*bill, 0)
	for _, row := range rows {
		tmpTime, err := time.Parse("2006/01/02 15:04:05", row[0])
		if err != nil {
			log.Fatal(err)
		}
		tmpAmmount, err := decimal.NewFromString(strings.TrimPrefix(row[5], "¥"))
		if err != nil {
			log.Fatal(err)
		}
		tmpBill := &bill{
			Timestamp: tmpTime,
			Commodity_type: row[1],
			Counterparty: row[2],
			Commodity: row[3],
			In_ex: row[4],
			Ammount: tmpAmmount,
			Pay_method: row[6],
			Order_number: row[8],
			Mer_number: row[9],
			Source: "微信",
			Remark: row[10],
		}
		bills = append(bills, tmpBill)
	}
	for _, b := range bills {
		res, err := stmt.Exec(b.Timestamp, b.In_ex, b.Commodity, "", b.Ammount, b.Pay_method, b.Counterparty, "", b.Order_number, b.Mer_number, b.Source, b.Remark)
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	}
}

func readAlipayBill() {

}