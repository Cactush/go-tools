1.nsqlookupd:
2.nsqd --lookupd-tcp-address=127.0.0.1:4160 
3.nsqadmin --lookupd-http-address=127.0.0.1:4161

go get github.com/nsqio/go-nsq
`https://github.com/nsqio/nsq/releases` 下载nsq的工具
放在 $GOPATH/bin 然后加入环境变量

注意
```cassandraql
json序列化数据的时候 内部数据成员必须大写,否则不能序列化
```
```cassandraql
https://blog.csdn.net/tian_lai_yuyuh/article/details/52700695
```