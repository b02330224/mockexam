# mockexam

项目需求

考场签到，名字丢入管道；
只有5个车道，最多供5个人同时考试；
考生按签到顺序依次考试，给予考生10%的违规几率；
每3秒钟巡视一次，发现违规的清出考场，否则输出考场时序良好；
所有考试者考完后，向MySQL数据库录入考试成绩；
成绩录入完毕通知考生，考生查阅自己的成绩；
当前目录下的成绩录入MySQL数据库,数据库允许一写多读；
再次查询成绩使用Redis缓存（二级缓存）；
整理优化代码，提高复用程度；
主要技术栈

管道并发
MySQL-Redis二级缓存
通用的数据库工具的封装
类库封装和复用