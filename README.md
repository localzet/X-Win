<p align="center"><a href="https://www.localzet.com" target="_blank">
  <img src="https://static.zorin.space/media/logos/ZorinProjectsSP.svg">
</a></p>

<p align="center">
  <a href="https://github.com/localzet/x-win/releases">
  <img src="https://img.shields.io/github/downloads/localzet/x-win/total.svg?label=%D0%A1%D0%BA%D0%B0%D1%87%D0%B8%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F" alt="Скачивания">
</a>
  <a href="https://github.com/localzet/x-win">
  <img src="https://img.shields.io/github/commit-activity/t/localzet/x-win?label=%D0%9A%D0%BE%D0%BC%D0%BC%D0%B8%D1%82%D1%8B" alt="Коммиты">
</a>
  <a href="https://github.com/localzet/x-win/releases/latest">
  <img src="https://img.shields.io/github/v/release/localzet/x-win?label=%D0%92%D0%B5%D1%80%D1%81%D0%B8%D1%8F" alt="Версия">
</a>
  <a href="https://github.com/localzet/x-win">
  <img src="https://img.shields.io/github/license/localzet/x-win?label=%D0%9B%D0%B8%D1%86%D0%B5%D0%BD%D0%B7%D0%B8%D1%8F" alt="Лицензия">
</a>
</p>

# X-WIN - Windows XRay Client

> Клиент для ядра XRay

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
