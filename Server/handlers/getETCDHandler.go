package handlers

// import (
// 	"errors"
// 	"strings"

// 	//"encoding/json"
// 	// "etcd_test/src/message"
// 	"fmt"
// 	"time"

// 	"github.com/golang/protobuf/proto"
// 	clientv3 "go.etcd.io/etcd/client/v3"
// 	"golang.org/x/net/context"
// )

// // 连接etcd集群
// func connectToEtcd(args ...string) (*clientv3.Client, error) {
// 	if len(args) == 0 {
// 		return nil, errors.New("ip or port is nil")
// 	}

// 	var endpoints []string
// 	for _, arg := range args {
// 		endpoints = append(endpoints, arg)
// 	}

// 	// 连接etcd集群
// 	cli, err := clientv3.New(clientv3.Config{
// 		Endpoints:   endpoints,
// 		DialTimeout: 5 * time.Second,
// 	})
// 	if err != nil {
// 		return nil, errors.New("connect etcd fail")
// 	}
// 	return cli, nil
// }

// // 初始化元数据
// func initMetadata() (map[string]*message.Metadata, error) {
// 	// 初始化一个切片用于存储元数据
// 	metadata := make(map[string]*message.Metadata)

// 	// 初始化meta_tables实例
// 	data1 := &message.Metadata{Name: "dist_by_mm", ShardCount: 12}
// 	keyName1 := "meta_table"
// 	metadata[keyName1] = data1

// 	data2 := &message.Metadata{Name: "dist_by_mm", ShardCount: 14}
// 	keyName2 := "meta_table1"
// 	metadata[keyName2] = data2

// 	if len(metadata) == 0 {
// 		return nil, errors.New("metadata init error")
// 	}
// 	return metadata, nil
// }

// // 将元数据序列化为protobuf并存储到etcd(单条循环写入)
// func putMetadataToEtcdSingle(cli *clientv3.Client, metadata map[string]*message.Metadata) error {
// 	if len(metadata) == 0 {
// 		return errors.New("metadata list is nil")
// 	}

// 	// 遍历元数据列表，逐个写入etcd
// 	for k, v := range metadata {
// 		// 将元数据序列化为protobuf
// 		protoData, err := proto.Marshal(v)
// 		if err != nil {
// 			return errors.New("metadata to protobuf fail")
// 		}

// 		// 存储protobuf数据到etcd
// 		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 		defer cancel()

// 		_, err = cli.Put(ctx, k, string(protoData))
// 		if err != nil {
// 			return errors.New("metadata write fail")
// 		}
// 	}
// 	return nil
// }

// // 将元数据序列化为protobuf并存储到etcd(批量写入)
// func putMetadataToEtcdBatch(cli *clientv3.Client, metadata map[string]*message.Metadata) error {
// 	// 创建事务
// 	txn := cli.Txn(context.Background())

// 	meta := make(map[string]string)
// 	for k, v := range metadata {
// 		// 将元数据序列化为protobuf
// 		protoData, err := proto.Marshal(v)
// 		if err != nil {
// 			return errors.New("metadata to protobuf fail")
// 		}
// 		meta[k] = string(protoData)
// 	}

// 	// 组装多个写入操作
// 	txn.If(clientv3.Compare(clientv3.CreateRevision("meta_table"), "=", 0)).
// 		// 此处不允许相同的键重复写入
// 		Then(
// 			clientv3.OpPut("meta_table", meta["meta_table"]),
// 			clientv3.OpPut("meta_table1", meta["meta_table1"]),
// 		)

// 	// 提交事务
// 	resp, err := txn.Commit()
// 	if err != nil {
// 		// handle error!
// 		fmt.Printf("txn commit failed, err: %v\n", err)
// 		return nil
// 	}

// 	// 检查事务结果
// 	if !resp.Succeeded {
// 		fmt.Println("txn failed, some keys already exist")
// 	} else {
// 		fmt.Println("txn succeeded")
// 	}
// 	return nil
// }

// func getMetadataFromEtcd(cli *clientv3.Client, metadataMap map[string]*message.Metadata, args ...interface{}) error {
// 	if len(args) == 0 {
// 		return errors.New("not key need read")
// 	}
// 	// 遍历key并从etcd中读取对应的key-val
// 	for _, arg := range args {
// 		switch v := arg.(type) {
// 		case map[string]*message.Metadata: // 以map的形式指定key(即读取刚写入的数据)
// 			{
// 				// map为空
// 				if len(v) == 0 {
// 					return errors.New("key map is nil")
// 				} else {
// 					for k, _ := range v {
// 						ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 						meta, err := cli.Get(ctx, k)
// 						cancel()
// 						if err != nil {
// 							return errors.New("read matadata fail")
// 						}
// 						if len(meta.Kvs) == 0 {
// 							return errors.New("no key-value found for the given key")
// 						}
// 						// 把protobuf解析为Metadata
// 						var retrievedMetadata message.Metadata
// 						err = proto.Unmarshal(meta.Kvs[0].Value, &retrievedMetadata)
// 						if err != nil {
// 							return errors.New("protobuf to Metadata fail")
// 						}
// 						metadataMap[k] = &retrievedMetadata
// 					}
// 				}
// 				break
// 			}
// 		case string: // 单个读或批量读
// 			{
// 				if strings.Contains(v, "prefix") { // prefix读
// 					str := strings.Split(v, ":")
// 					key := str[0]

// 					// 创建事务
// 					txn := cli.Txn(context.Background())

// 					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 					meta, err := cli.Get(ctx, key, clientv3.WithPrefix())
// 					cancel()
// 					if err != nil {
// 						return errors.New("read matadata error")
// 					}
// 					if len(meta.Kvs) == 0 {
// 						return errors.New("no key-value found for the given key")
// 					}
// 					for _, m := range meta.Kvs {
// 						// 把protobuf解析为Metadata
// 						var val message.Metadata
// 						err = proto.Unmarshal(m.Value, &val)
// 						if err != nil {
// 							return errors.New("protobuf to Metadata fail")
// 						}
// 						metadataMap[string(m.Key)] = &val
// 					}
// 					// 提交事务
// 					resp, err := txn.Commit()
// 					if err != nil {
// 						return errors.New("txn commit failed")
// 					}

// 					// 检查事务结果
// 					if !resp.Succeeded {
// 						return errors.New("txn failed, some keys already exist")
// 					} else {
// 						return errors.New("txn succeeded")
// 					}
// 				} else { // 单条读
// 					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 					meta, err := cli.Get(ctx, v)
// 					cancel()
// 					if err != nil {
// 						return errors.New("read matadata error")
// 					}
// 					if len(meta.Kvs) == 0 {
// 						return errors.New("no key-value found for the given key")
// 					}
// 					var val message.Metadata
// 					err = proto.Unmarshal(meta.Kvs[0].Value, &val)
// 					if err != nil {
// 						return errors.New("protobuf to Metadata fail")
// 					}
// 					metadataMap[string(meta.Kvs[0].Key)] = &val
// 				}
// 				break
// 			}
// 		default:
// 			return errors.New("args is unknown data type")
// 		}
// 	}
// 	return nil
// }

// // 监听etcd上指定key的变化，如果发生变化则更新缓存中的元数据
// func watchEtcdAndUpdate(cli *clientv3.Client, metadataMap map[string]*message.Metadata, key string) error {
// 	if key == "" {
// 		return errors.New("not key need watch")
// 	}

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	watch := cli.Watch(ctx, key, clientv3.WithPrefix(), clientv3.WithPrevKV())
// 	for w := range watch {
// 		for _, ev := range w.Events {
// 			fmt.Printf("Type: %s Key: %s Value: %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)

// 			// 根据变化修改缓存中的元数据
// 			switch ev.Type {
// 			case clientv3.EventTypePut:
// 				{
// 					var val message.Metadata
// 					err := proto.Unmarshal(ev.Kv.Value, &val)
// 					if err != nil {
// 						return errors.New("protobuf to Metadata fail")
// 					}
// 					metadataMap[string(ev.Kv.Key)] = &val

// 					for k, v := range metadataMap {
// 						fmt.Println("tableName:", k)
// 						fmt.Println("metadata ShardCount:", v.GetShardCount())
// 					}
// 					break
// 				}
// 			case clientv3.EventTypeDelete:
// 				{
// 					if ev.PrevKv.Key != nil {
// 						delete(metadataMap, string(ev.Kv.Key))
// 					}
// 					break
// 				}
// 			default:
// 				return errors.New("watch error")
// 			}
// 		}
// 	}
// 	return nil
// }

// func getETCDHandlers() {
// 	// 初始化一个缓存，模拟其它kingproxy节点
// 	metadataMap := make(map[string]*message.Metadata)

// 	// 连接etcd集群
// 	cli, err := connectToEtcd("104.168.125.34:2379") //("120.92.144.250:2379")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer cli.Close()

// 	// 初始化元数据
// 	metadata, err := initMetadata()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// 写入元数据
// 	//err = putMetadataToEtcdSingle(cli, metadata)
// 	//if err != nil {
// 	// fmt.Println(err)
// 	// return
// 	//}

// 	// 批量写入元数据
// 	err = putMetadataToEtcdBatch(cli, metadata)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// 读取元数据
// 	getMetadataFromEtcd(cli, metadataMap, "meta_table:prefix")

// 	// pre result
// 	for k, v := range metadataMap {
// 		fmt.Println("tableName:", k)
// 		fmt.Println("metadata ShardCount:", v.GetShardCount())
// 	}

// 	fmt.Println("--------------------------------------------------------------------------------")

// 	// watch key
// 	//watchEtcdAndUpdate(cli, metadataMap, "meta_table")

// 	return
// }
