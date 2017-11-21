package mysql


import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "log"


var DB *sql.DB



func Init() {
    var err error
	DB, err = sql.Open("mysql", "root:root@/web")
    if err !=nil {
		fmt.Printf("ERROR")
	}
	
}

func Insert(db *sql.DB,name string, passwd string,info string) bool {

	 _, err := db.Exec(`INSERT user (username,user_passwd,user_info) values (?,?,?)`,name,passwd,info)
	if err != nil {
		fmt.Println(err)
		return false
	}
	
	return true
}

func Query(db *sql.DB, name string) (passwd string,info string,flag bool) {

	rows, err := db.Query("select user_passwd, user_info from user where username = ?", name)
	if err != nil {
		log.Println(err)
		return passwd,info,flag
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&passwd, &info)
		if err != nil {
			log.Print(err)
			return passwd,info,flag
		}
	}
	flag = true 
	return passwd,info,flag
}



