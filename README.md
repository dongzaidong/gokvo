# gokvo
Implement message notifications based on notifications in ios

# Introduction
项目开发中，经常遇到A业务完成需要通知B业务或者C业务等等去做一些响应，为了完成这种项目中不同位置的信息发送，使用键值监听通知模式，但是在go没有官方实现，所以作者根据以前ios开发经验，模仿了notification机制做了一个适用于go的框架

# How to install
```
go get -u https://github.com/dongzaidong/gokvo
```

# How to use
```
	// A 业务
	var obser = struct{}{}
	gokvo.AddObserver(obser, "send", func(a interface{}) {
		fmt.Println("监听到值：", a)
	})

	
	// B 业务
	err := 	gokvo.Post("send", "hello world")
	fmt.Println(err)

	// 移除监听
	gokvo.RemoveObserver(obser, "send")
```

