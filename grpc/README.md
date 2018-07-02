# grpc sample

## 启动server端
```$xslt
cd ${GOPATH}/src/github.com/zhongqiuwood
git clone https://github.com/zhongqiuwood/sample.git
cd sample/grpc/server
./run.sh

2018/07/02 08:00:34 listen on: 127.0.0.1:10001

```

## 启动client端
```$xslt
cd ${GOPATH}/src/github.com/zhongqiuwood/sample/grpc/client
./run.sh

2018/07/02 08:01:55 Try to connect to: 127.0.0.1:10001
2018/07/02 08:01:55 Initiating Chat with peer address: 127.0.0.1:10001
2018/07/02 08:01:55 Established Chat with peer address: 127.0.0.1:10001
2018/07/02 08:01:55 Current context deadline = 0001-01-01 00:00:00 +0000 UTC, ok = false, initiatedStream=true
2018/07/02 08:01:55 newHelloMessage: ID:"peer0" address:"127.0.0.1:10000" 
2018/07/02 08:01:55 Received DISC_HELLO, parsing out Peer identification
2018/07/02 08:01:55 Received DISC_HELLO from endpoint=endpoint:<ID:"peer1" address:"127.0.0.1:10001" > 
2018/07/02 08:01:55 registered handler with key: peer1
2018/07/02 08:01:55 The Peer's bi-directional stream to ID:"peer1" address:"127.0.0.1:10001"  is established, from event DISC_HELLO

```