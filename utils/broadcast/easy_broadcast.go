package broadcast

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
	"github.com/szpnygo/mab/utils/snow"
)

/*
轻量级广播系统，使用Redis订阅模式
在没有性能要求、数据持久化、非重要场景下使用
*/
type EasyBroadcast struct {
	name        string //订阅的redis频道名
	redisClient *redis.Client
	sf          *snowflake.Node
	clients     sync.Map //储存维护client
}

type Event struct {
	Topic string
	Data  []byte
}

type Client struct {
	name      string
	topics    []string
	c         chan *Event
	broadcast *EasyBroadcast
}

func NewEasyBroadcast(name string, redisAddr string) (*EasyBroadcast, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", //有需要的时候再考虑完整的redis配置
		DB:       2,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	sf, err := snow.NewSnowNode()
	if err != nil {
		return nil, err
	}

	return &EasyBroadcast{
		name:        name,
		redisClient: rdb,
		sf:          sf,
	}, nil
}

// Send 发送消息
func (eb *EasyBroadcast) Send(ctx context.Context, topic string, msg interface{}) error {
	body, _ := json.Marshal(msg)
	data, _ := json.Marshal(Event{
		Topic: topic,
		Data:  body,
	})
	_, err := eb.redisClient.Publish(ctx, eb.name, string(data)).Result()

	return err
}

// NewClient 创建一个新的客户端
func (eb *EasyBroadcast) NewClient(topics ...string) (*Client, func()) {
	client := &Client{
		name:      "client_" + eb.sf.Generate().String(),
		topics:    topics,
		c:         make(chan *Event),
		broadcast: eb,
	}

	eb.clients.Store(client.name, client)

	return client, func() {
		eb.clients.Delete(client.name)
		close(client.c)
	}
}

// Start 开始监听
func (eb *EasyBroadcast) Start(ctx context.Context) {
	go func() {
		subscribe := eb.redisClient.Subscribe(ctx, eb.name)
		for {
			msg, err := subscribe.ReceiveMessage(ctx)
			if err == nil {
				var event Event
				err := json.Unmarshal([]byte(msg.Payload), &event)
				if err == nil {
					eb.broadcast(&event)
				} else {
					fmt.Println(err.Error())
				}
			} else {
				fmt.Println(err.Error())
			}
		}
	}()
}

//广播消息
func (eb *EasyBroadcast) broadcast(event *Event) {
	eb.clients.Range(func(key, value interface{}) bool {
		client := value.(*Client)
		client.send(event)
		return true
	})
}

func (client *Client) send(event *Event) {
	for _, topic := range client.topics {
		if topic == event.Topic {
			client.c <- event
			break
		}
	}
}

// Receive 接受消息
func (client *Client) Receive() (string, []byte, error) {
	event, ok := <-client.c
	if ok {
		return event.Topic, event.Data, nil
	}

	return "", nil, fmt.Errorf("client closed")
}
