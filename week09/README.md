# Q&A

1. 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用。
   fix length: 每个包发送固定长度， 非常不灵活。
   delimiter based: 可以在数据包之间设置边界，如添加特殊符号，这样，接收端通过这个边界就可以将不同的数据包拆分开， 特殊符号的选取会是问题，而且不通用。
   length field based: 每次发送一个应用数据包前在前面加上四个字节的包长度值，指明这个应用包的真实长度。

2. 实现一个从 socket connection 中解码出 goim 协议的解码器。
   参见 main.go
