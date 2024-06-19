# QR Code Generator Microservice

## Overview (Umumiy koʻrinish)

Ushbu mikroservis user so'rovlaridan `QR` kodlarini ishlab chiqaradi. U `POST` so'rovi orqali ma'lumotlarni qabul qiladi, taqdim etilgan ma'lumotlarni `QR` kodini rasm sifatida qaytaradi.

Misol:
```shell
curl -X POST \
    --form "size=256" \
    --form "content=https://example.com" \
    --output data/qrcode.png \
    http://localhost:8080/generate
```



