# vpn-overrides
Дисклеймер: я просто скопировал весь wireguard-go сюда, потому что лень. Не судите строго

Вся суть находится в папке overrides. В этой же папке нужно создать файл config.json, который должен иметь следующий вид:

{
  "tunAddr": <Ваш локальный адрес для туннеля>,
  "dnsAddr": "8.8.8.8",
  "privateKey": <Ваш приватный ключ>,
  "publicKey": <Публичный ключ VPN-сервера>,
  "serverAddr": <Адрес VPN-сервера>,
  "serverPort": <Порт VPN-сервера>,
  "allowedIp": "0.0.0.0/0",
  "network": "tcp",
  "address": <Конечный адрес>
}

P.S. network и address пока не испольщуются. Это остается автору как TODO