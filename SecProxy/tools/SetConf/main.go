package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

const (
	EtcdKey = "/oldboy/backend/secskill/product"
)

// SecInfoConf 秒杀商品相关信息的结构体
type SecInfoConf struct {
	ProductId int //商品的ID
	StartTime int // 秒杀开始时间
	EndTime   int // 秒杀结束时间
	Status    int // 秒杀商品的状态
	Total     int // 商品的总数
	Left      int // 剩余商品数量
}

func SetLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	var SecInfoConfArr []SecInfoConf
	SecInfoConfArr = append(
		SecInfoConfArr,
		SecInfoConf{
			ProductId: 1029,
			StartTime: 1505008800,
			EndTime:   1505012400,
			Status:    0,
			Total:     1000,
			Left:      1000,
		},
	)
	SecInfoConfArr = append(
		SecInfoConfArr,
		SecInfoConf{
			ProductId: 1027,
			StartTime: 1505008800,
			EndTime:   1505012400,
			Status:    0,
			Total:     2000,
			Left:      1000,
		},
	)

	data, err := json.Marshal(SecInfoConfArr)
	if err != nil {
		fmt.Println("json failed, ", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//cli.Delete(ctx, EtcdKey)
	//return
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func main() {
	SetLogConfToEtcd()
}
