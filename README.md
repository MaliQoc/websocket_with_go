# websocket_with_go

 Sunucu, WebSocket protokolü aracılığıyla bağlantıları kabul eder ve alınan mesajları geri gönderir. Testler, sunucunun doğru bir şekilde çalıştığını ve mesaj alışverişini doğrular.

## Kurulum ve Çalıştırma (Visual Studio Code)

1. Gerekli kütüphaneleri yükleyin:
 
    ```bash
    go get github.com/gorilla/websocket
    ```
(Go ile basit bir WebSocket sunucusu yazmak için genellikle github.com/gorilla/websocket kütüphanesini kullanırız.)

2. Websocketi çalıştırın:

    ```bash
    go run main.go
    ```
    
3. Test edin:

   ```bash
    go test -v
    ```
