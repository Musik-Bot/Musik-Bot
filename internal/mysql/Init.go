package mysql

func InitDatabase() {
	conn := GetConn()
	stmt, err := conn.Prepare("CREATE TABLE IF NOT EXISTS `downloads` (`ID` int(11) unsigned NOT NULL auto_increment, `url` varchar(512), `FileName` varchar(512), `Name` varchar(512), PRIMARY KEY  (`ID`));")
	if err != nil {
		panic(err)
	}
	stmt.Exec()
	defer conn.Close()
}
