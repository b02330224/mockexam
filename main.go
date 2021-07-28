package main

import (
	"time"
	"fmt"
	"mockexam/utils"
	"sync"
)

/*
考场签到，名字丢入管道；
只有5个车道，最多供5个人同时考试；
考生按签到顺序依次考试，给予考生10%的违规几率；
每3秒钟巡视一次，发现违规的清出考场，否则输出考场时序良好；
所有考试者考完后，向MySQL数据库录入考试成绩；
成绩录入完毕通知考生，考生查阅自己的成绩；
当前目录下的成绩录入MySQL数据库,数据库允许一写多读；
再次查询成绩使用Redis缓存（二级缓存）；
整理优化代码，提高复用程度；
*/

var (
	wg sync.WaitGroup
)

/*主程序*/
func main() {
	for i := 0; i < 20; i++ {
		chNames <- utils.GetRandomName()
	}
	close(chNames)

	/*巡考*/
	go Patrol()

	/*考生并发考试*/
	for name := range chNames {
		wg.Add(1)
		go func(name string) {
			TakeExam(name)
			wg.Done()
		}(name)
	}

	wg.Wait()
	fmt.Println("考试完毕！")

	/*录入成绩*/
	wg.Add(1)
	go func() {
		utils.WriteScore2Mysql(scoreMap)
		wg.Done()
	}()
	//故意给一个时间间隔，确保WriteScore2DB先抢到数据库的读写锁
	<-time.After(1 * time.Second)

	/*考生查询成绩*/
	for _, name := range examers {
		wg.Add(1)
		go func(name string) {
			QueryScore(name)
			wg.Done()
		}(name)
	}
	<-time.After(1 * time.Second)
	for _, name := range examers {
		wg.Add(1)
		go func(name string) {
			QueryScore(name)
			wg.Done()
		}(name)
	}

	wg.Wait()
	fmt.Println("END")
}
