package main

import (  
	// "database/sql"
	"log"
	"time"
	"fmt"
	"strings"

	// "github.com/xuri/excelize/v2"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
    "github.com/xuri/excelize/v2"
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
	readWeiChatBill()
	readAlipayBill()
}

func readWeiChatBill() {
 //   // 准备数据库连接
	// db, err := sql.Open("mysql", "root:Ltf*ms93@tcp(127.0.0.1:3306)/bill?parseTime=true")  
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	
	// // 测试数据库连接
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }


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
		// for _, colCell := range row {
		// 	fmt.Print(colCell, "\t")
		// }
		// fmt.Println()
		tmpTime, err := time.Parse("2006/01/02 15:04:05", row[0])
		if err != nil {
			log.Fatal(err)
		}
		tmpAmmount, err := decimal.NewFromString(strings.TrimPrefix(row[5], "¥"))
		if err != nil {
			log.Fatal(err)
		}
		tmpBill := &bill{
			timestamp: tmpTime,
			commodity_type: row[1],
			counterparty: row[2],
			commodity: row[3],
			in_ex: row[4],
			ammount: tmpAmmount,
			pay_method: row[6],
			tran_account: "",
			order_number: row[8],
			mer_number: row[9],
			source: "微信",
			remark: row[10],
		}
		bills = append(bills, tmpBill)
		fmt.Println(tmpBill)
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
}

func readAlipayBill() {

}