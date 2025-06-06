# vcode

`vcode` — DevOps CLI-інструмент, написаний на Go з використанням `cobra`.

## Команди

- `vcode ping <host>` — перевіряє доступність хоста через TCP:80

## Запуск

```bash
go build -o vcode
./vcode ping google.com
