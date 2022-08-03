package main

import (
	"database/sql"
	"fmt"
	"fyne-app/file/hello"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}

	defer con.Close()

	ids := hello.Input("update ID")
	id, _ := strconv.Atoi(ids)

	qry := "select * from mydata where id = ?"

	rw := con.QueryRow(qry, id)
	tgt := mydatafmRw(rw)

	fmt.Println(tgt.Str())

	f := hello.Input("delete it? (y/n)")
	if f == "y" {
		qry = "delete from mydata where id=?"
		con.Exec(qry, id)
	}

	// ids := hello.Input("update ID")
	// id, _ := strconv.Atoi(ids)

	// qry := "select * from mydata where id = ?"

	// rw := con.QueryRow(qry, id)
	// tgt := mydatafmRw(rw)

	// ae := strconv.Itoa(tgt.Age)
	// nm := hello.Input("name(" + tgt.Name + ")")
	// ml := hello.Input("mail(" + tgt.Mail + ")")
	// ge := hello.Input("age(" + ae + ")")

	// ag, _ := strconv.Atoi(ge)

	// if nm == "" {
	// 	nm = tgt.Name
	// }

	// if ml == "" {
	// 	ml = tgt.Mail
	// }

	// if ge == "" {
	// 	ag = tgt.Age
	// }

	// qry2 := "update mydata set name=?, mail=?, age=? where id=?"

	// con.Exec(qry2, nm, ml, ag, id)

	showRecode(con)

}

func showRecode(con *sql.DB) {
	qry := "select * from mydata"
	rs, _ := con.Query(qry)
	for rs.Next() {
		fmt.Println(mydatafmRws(rs))
	}
}

func mydatafmRws(rs *sql.Rows) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}

func mydatafmRw(rs *sql.Row) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}

//var qry string = "select * from mydata where name LIKE ? OR mail LIKE ?"

// func main() {
// 	con, er := sql.Open("sqlite3", "data.sqlite3")
// 	if er != nil {
// 		panic(er)
// 	}

// 	defer con.Close()

// 	for true {

// 		s := hello.Input("find")
// 		if s == "" {
// 			break
// 		}

// 		rs, er := con.Query(qry, "%"+s+"%", "%"+s+"%")
// 		if er != nil {
// 			panic(er)
// 		}

// 		for rs.Next() {
// 			var md Mydata

// 			er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
// 			if er != nil {
// 				panic(er)
// 			}
// 			fmt.Println(md.Str())
// 		}
// 	}

// 	fmt.Println("*** end ***")
// }

type Mydata struct {
	ID   int
	Name string
	Mail string
	Age  int
}

func (m *Mydata) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\" " + m.Mail + "," + strconv.Itoa(m.Age) + ">"
}
