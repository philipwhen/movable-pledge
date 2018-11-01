package utils

import (
	"database/sql"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3" //sqlite3 dirver
	"github.com/sifbc/pledge/apiserver/define"
)

type appContext struct {
	db *sql.DB
}

/* ##############################################
* description: 连接数据库
* input:       driverName	用的驱动名称
*			   dbName		数据库名称
* output:      返回数据库内容
* ###############################################*/
func ConnectDB(driverName string, dbName string) (*appContext, error) {
	db, err := sql.Open(driverName, dbName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &appContext{db}, nil
}

/* ##############################################
* description: 断开与数据库的链接
* input:       无
* output:      error信息
* ###############################################*/
func (c *appContext) DisConnectDB() error {
	err := c.db.Close()
	if err != nil {
		return err
	}
	return nil
}

/* ##############################################
* description: 创建数据表
* input:       无
* output:      error信息
* ###############################################*/
func (c *appContext) CreateTable() error {
	sqlStmt := `
	create table PledgeWarningInfo(
	    PledgeNoStorage text not null primary key,
	    ChannelSeq text,
	    PledgeName text,
	    SocialCreditCode text,
	    CoreEnterpriseName text,
	    WarningMsg text,
		DataType integer
	);
	`
	// delete from foo;
	_, err := c.db.Exec(sqlStmt)
	if err != nil {
		Log.Errorf("Sqlite3 create table err: %q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}

/* ##############################################
* description: 查询数据并更新增量数据
* input:       PledgeNoStorage:    质押物编号
*              ChannelSeq:         渠道编号
*              PledgeName:         质押物名称
*              SocialCreditCode:   统一社会信用代码
*              CoreEnterpriseName: 仓储企业名称
*              WarningMsg:         告警信息
* output:      flag:               数据库中是否已存有质押物信息
*              error:              信息
* ###############################################*/
func (c *appContext) QueryAndUpdateTable(info define.PledgeWarningInfo) (bool, error) {
	flag, savedInfo, err := c.QueryTable(info.PledgeNoStorage)
	if err != nil {
		Log.Errorf("query table err, %s", err.Error())
		return flag, err
	}
	if !flag {
		err := c.UpdateTable(info.PledgeNoStorage, savedInfo.WarningMsg+info.WarningMsg)
		if err != nil {
			Log.Errorf("update table err, %s", err.Error())
			return flag, err
		}
	}
	return flag, nil
}

/* ##############################################
* description: 插入数据
* input:       PledgeNoStorage:    质押物编号
*              ChannelSeq:         渠道编号
*              PledgeName:         质押物名称
*              SocialCreditCode:   统一社会信用代码
*              CoreEnterpriseName: 仓储企业名称
*              WarningMsg:         告警信息
* output:      error:              信息
* ###############################################*/
func (c *appContext) InsertTable(info define.PledgeWarningInfo) error {
	tx, err := c.db.Begin()
	if err != nil {
		Log.Errorf("Sqlite3 insert table begin error , %s", err.Error())
		return err
	}
	stmt, err := c.db.Prepare("insert into PledgeWarningInfo(PledgeNoStorage, ChannelSeq, PledgeName, SocialCreditCode, CoreEnterpriseName, WarningMsg, DataType) values(?,?,?,?,?,?,?)")
	if err != nil {
		Log.Errorf("Sqlite3 insert table prepare error , %s", err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(info.PledgeNoStorage, info.ChannelSeq, info.PledgeName, info.SocialCreditCode, info.CoreEnterpriseName, info.WarningMsg, info.DataType)
	if err != nil {
		Log.Errorf("Sqlite3 insert table error , %s", err.Error())
		return err
	}
	err = tx.Commit()
	return err
}

/* ##############################################
* description: 根据质押物编号读取数据
* input:       PledgeNoStorage:    质押物编号
* output:      flag:               数据库中是否未存有质押物信息
*              PledgeNoStorage:    质押物编号
*              ChannelSeq:         渠道编号
*              PledgeName:         质押物名称
*              SocialCreditCode:   统一社会信用代码
*              CoreEnterpriseName: 仓储企业名称
*              WarningMsg:         告警信息
*			   error:              信息
* ###############################################*/
func (c *appContext) QueryTable(pledgeNoStorage string) (bool, define.PledgeWarningInfo, error) {
	var info define.PledgeWarningInfo
	rows := c.db.QueryRow("select * from PledgeWarningInfo where PledgeNoStorage = ?", pledgeNoStorage)
	err := rows.Scan(&info.PledgeNoStorage, &info.ChannelSeq, &info.PledgeName, &info.SocialCreditCode, &info.CoreEnterpriseName, &info.WarningMsg, &info.DataType)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			Log.Errorf("%s\n", err.Error())
			return true, info, nil
		} else {
			Log.Errorf("Sqlite3 query table error , %s", err.Error())
			return false, info, err
		}
	}
	return false, info, nil
}

/* ##############################################
* description: 读取数据表
* input:       无
* output:      PledgeNoStorage:    质押物编号
*              ChannelSeq:         渠道编号
*              PledgeName:         质押物名称
*              SocialCreditCode:   统一社会信用代码
*              CoreEnterpriseName: 仓储企业名称
*              WarningMsg:         告警信息
*			   error:              信息
* ###############################################*/
func (c *appContext) QueryAllTable() ([]define.PledgeWarningInfo, error) {
	rows, err := c.db.Query("select * from PledgeWarningInfo")
	if err != nil {
		Log.Errorf("Sqlite3 query table error , %s", err.Error())
		return nil, err
	}
	defer rows.Close()
	infoList := []define.PledgeWarningInfo{}
	for rows.Next() {
		info := new(define.PledgeWarningInfo)
		err = rows.Scan(&info.PledgeNoStorage, &info.ChannelSeq, &info.PledgeName, &info.SocialCreditCode, &info.CoreEnterpriseName, &info.WarningMsg, &info.DataType)
		if err != nil {
			Log.Errorf("Sqlite3 query table Scan error , %s", err.Error())
			return nil, err
		}
		infoList = append(infoList, *info)
	}
	err = rows.Err()
	if err != nil {
		Log.Errorf("Sqlite3 query table rows error , %s", err.Error())
	}
	return infoList, nil
}

/* ##############################################
* description: 更新质押物的告警信息
* input:       PledgeNoStorage:    质押物编号
*			   WarningMsg:         告警信息
* output:      error:              信息
* ###############################################*/
func (c *appContext) UpdateTable(pledgeNoStorage string, msg string) error {
	stmt, err := c.db.Prepare("update PledgeWarningInfo set WarningMsg = ? WHERE PledgeNoStorage = ?")
	if err != nil {
		Log.Errorf("Sqlite3 update table prepare error , %s", err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(msg, pledgeNoStorage)
	if err != nil {
		Log.Errorf("Sqlite3 update table exec error , %s", err.Error())
		return err
	}
	return nil
}

/* ##############################################
* description: 删除质押物信息
* input:       PledgeNoStorage:    质押物编号
* output:      error:              信息
* ###############################################*/
func (c *appContext) DeleteTable(pledgeNoStorage string) error {
	stmt, err := c.db.Prepare("delete from PledgeWarningInfo where PledgeNoStorage = ?")
	if err != nil {
		Log.Errorf("Sqlite3 delete table prepare error , %s", err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(pledgeNoStorage)
	if err != nil {
		Log.Errorf("Sqlite3 delete table exec error , %s", err.Error())
		return err
	}
	return nil
}

/* ##############################################
* description: 判断文件或文件夹是否存在
* input:       文件或文件夹
* output:      没有error时，true表示文件存在，false
*			   表示文件不存在。
* ###############################################*/
func FileOrDirectoryExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil //文件或文件夹存在
	}
	if os.IsNotExist(err) {
		return false, nil //文件或文件夹不存在
	}
	return true, err //其他错误
}
