package main

import (
	"fmt"

	_ "github.com/localzet/xray-wrapper/xray"
)

func main() {
	fmt.Println("XRay-wrapper")
	fmt.Println("Сreated by localzet")
	fmt.Println("https://github.com/localzet")

	// Для тестирования можно использовать функцию xray.RunTest(path, port, logLevel, logPath isSocks, isUdpEnabled).
	// раскомментируйте эти строки и замените переменную "path" на ваш путь конфигурации,
	// и замените переменную "port" на порт, который вы хотите прослушивать.
	// Для доступа к информации журнала установите "logLevel" и "logPath"
	// Если вы хотите запустить его в режиме socks, установите "isSocks" на true, иначе false.
	// Если вы хотите проксировать трафик UDP, установите "isUdpEnabled" на true, иначе false. (только для socks)

	// path := "Например: C:/config.json"
	// port := "Например: 10801"
	// logLevel := "none, debug, info, warning или error"
	// logPath := "Например: C:/Log"
	// isSocks := "Например: true"
	// isUdpEnabled := "Например: true"
	// xray.RunTest(path, port, logLevel, logPath, isSocks, isUdpEnabled)
}
