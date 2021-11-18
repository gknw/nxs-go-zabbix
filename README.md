# nxs-go-zabbix

This Go package provides access to Zabbix API v5.0.
Also see older versions in other branches.

At the time not all Zabbix API methods are implemented, but work in progress.

## Install

```
go get github.com/nixys/nxs-go-zabbix/v5
```

## Example of usage

*You may find more examples in unit-tests in this repository*

**Get all hosts on Zabbix server:**

```go
package main

import (
	"fmt"
	"os"

	"github.com/nixys/nxs-go-zabbix/v5"
)

func zabbixLogin(z *zabbix.Context, zbxHost, zbxUsername, zbxPassword string) {

	if err := z.Login(zbxHost, zbxUsername, zbxPassword); err != nil {
		fmt.Println("Login error:", err)
		os.Exit(1)
	} else {
		fmt.Println("Login: success")
	}
}

func zabbixLogout(z *zabbix.Context) {

	if err := z.Logout(); err != nil {
		fmt.Println("Logout error:", err)
		os.Exit(1)
	} else {
		fmt.Println("Logout: success")
	}
}

func main() {

	var z zabbix.Context

	/* Get variables from environment to login to Zabbix server */
	zbxHost := os.Getenv("ZABBIX_HOST")
	zbxUsername := os.Getenv("ZABBIX_USERNAME")
	zbxPassword := os.Getenv("ZABBIX_PASSWORD")
	if zbxHost == "" || zbxUsername == "" || zbxPassword == "" {
		fmt.Println("Login error: make sure environment variables `ZABBIX_HOST`, `ZABBIX_USERNAME` and `ZABBIX_PASSWORD` are defined")
		os.Exit(1)
	}

	/* Login to Zabbix server */
	zabbixLogin(&z, zbxHost, zbxUsername, zbxPassword)
	defer zabbixLogout(&z)

	/* Get all hosts */
	hObjects, _, err := z.HostGet(zabbix.HostGetParams{
		GetParameters: zabbix.GetParameters{
			Output: zabbix.SelectExtendedOutput,
		},
	})
	if err != nil {
		fmt.Println("Hosts get error:", err)
		return
	}

	/* Print names of retrieved hosts */
	fmt.Println("Hosts list:")
	for _, h := range hObjects {
		fmt.Println("-", h.Host)
	}
}
```

Run:

```
ZABBIX_HOST="https://zabbix.yourdomain.com/api_jsonrpc.php" ZABBIX_USERNAME="Admin" ZABBIX_PASSWORD="PASSWORD" go run main.go
```

## Testing

Read prepared docker-compose file:

```
cat docker/development.yaml
```

Start Zabbix docker containers:

```
docker-compose -f docker/development.yaml up -d
```

Run test:

```
ZABBIX_HOST=http://127.0.0.1:8080/api_jsonrpc.php ZABBIX_USERNAME=Admin ZABBIX_PASSWORD=zabbix go test -v
```
