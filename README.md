<p align="center"><a href="https://www.localzet.com" target="_blank">
  <img src="https://static.zorin.space/media/logos/LocalzetGroup.png" width="400">
</a></p>

# X-WIN - Windows XRay Client

> Клиент для ядра XRay

[![Downloads](https://img.shields.io/github/downloads/localzet/x-win/total.svg?label=downloads%20%28total%29)](https://github.com/localzet/x-win/releases)
[![DownloadsLatest](https://img.shields.io/github/downloads/localzet/x-win/latest/total?label=downloads%20%28latest%29)](https://github.com/localzet/x-win/releases/latest)
[![LatestVersion](https://img.shields.io/github/v/release/localzet/x-win?label=latest%20version)](https://github.com/localzet/x-win/releases/latest)
[![CodeFactor](https://www.codefactor.io/repository/github/localzet/x-win/badge)](https://www.codefactor.io/repository/github/localzet/x-win)

X-WIN — это клиент, поддерживающий [XRay](https://github.com/XTLS/Xray-core). Он предоставляет простой в использовании интерфейс для настройки и управления прокси-серверами и позволяет пользователям переключаться между различными серверами.

## Разработка

- Если вы новичок, загрузите приложение с [releases](https://github.com/localzet/x-win/releases/latest).

- Но если вы хотите получить исходный код клиента, выполните следующие действия:
- Загрузите [requirements](#requirements)
- Клонируйте копию репозитория:
  ```
  git clone "https://github.com/localzet/x-win.git"
  ```
  - Перейдите в каталог:
  ```
  cd x-win
  ```
  - Создайте файл `XRayCore.dll` и скопируйте его в каталог `/src/Libraries`:
  ```
  cd xray
  go build --buildmode=c-shared -o XRayCore.dll -trimpath -ldflags "-s -w -buildid=" .
  md ..\src\Libraries
  copy XRayCore.dll ..\src\Libraries
  ```

- Загрузите службу `X-TUN` (для вашей ОС) по [этой ссылке](https://github.com/localzet/x-tun/releases/latest), извлеките и скопируйте в каталог `/src/TUN`.

- Загрузите файлы `geoip.dat` и `geosite.dat` и скопируйте их в каталог `/src`:
  ```
  cd ..\src
  curl https://github.com/v2fly/geoip/releases/latest/download/geoip.dat -o geoip.dat -L
  curl https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat -o geosite.dat -L
  ```

- Запустите проект:
  ```
  dotnet run
  ```

## Требования

- Go https://go.dev/dl/
- .Net https://dotnet.microsoft.com/download
