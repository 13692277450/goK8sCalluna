package handlers

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdManager struct {
	client *clientv3.Client
}

// NewEtcdManager 创建 etcd 管理器
func NewEtcdManager(endpoints []string) (*EtcdManager, error) {
	cfg := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create etcd client: %v", err)
	}

	return &EtcdManager{client: client}, nil
}

// Close 关闭客户端连接
func (em *EtcdManager) Close() error {
	return em.client.Close()
}

// GetAllKeys 获取所有键值对
func (em *EtcdManager) GetAllKeys(ctx context.Context, prefix string) (map[string]string, error) {
	resp, err := em.client.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("failed to get keys: %v", err)
	}

	result := make(map[string]string)
	for _, kv := range resp.Kvs {
		result[string(kv.Key)] = string(kv.Value)
	}

	return result, nil
}

// ListKeys 列出指定前缀的所有键
func (em *EtcdManager) ListKeys(ctx context.Context, prefix string) ([]string, error) {
	resp, err := em.client.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return nil, fmt.Errorf("failed to list keys: %v", err)
	}

	var keys []string
	for _, kv := range resp.Kvs {
		keys = append(keys, string(kv.Key))
	}

	return keys, nil
}

// WatchPrefix 监视指定前缀的键变化
func (em *EtcdManager) WatchPrefix(ctx context.Context, prefix string) {
	watchChan := em.client.Watch(ctx, prefix, clientv3.WithPrefix())

	fmt.Printf("开始监视前缀: %s\n", prefix)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("停止监视")
			return
		case watchResp := <-watchChan:
			if watchResp.Err() != nil {
				log.Printf("监视错误: %v", watchResp.Err())
				continue
			}

			for _, event := range watchResp.Events {
				switch event.Type {
				case clientv3.EventTypePut:
					if event.IsCreate() {
						fmt.Printf("[CREATE] 键: %s, 值: %s\n",
							string(event.Kv.Key),
							string(event.Kv.Value))
					} else {
						fmt.Printf("[UPDATE] 键: %s, 值: %s\n",
							string(event.Kv.Key),
							string(event.Kv.Value))
					}
				case clientv3.EventTypeDelete:
					fmt.Printf("[DELETE] 键: %s\n", string(event.Kv.Key))
				}
			}
		}
	}
}

// WatchSpecificKey 监视特定键的变化
func (em *EtcdManager) WatchSpecificKey(ctx context.Context, key string) {
	watchChan := em.client.Watch(ctx, key)

	fmt.Printf("开始监视键: %s\n", key)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("停止监视")
			return
		case watchResp := <-watchChan:
			if watchResp.Err() != nil {
				log.Printf("监视错误: %v", watchResp.Err())
				continue
			}

			for _, event := range watchResp.Events {
				switch event.Type {
				case clientv3.EventTypePut:
					fmt.Printf("[键变化] %s: %s -> %s\n",
						string(event.Kv.Key),
						"新值", string(event.Kv.Value))
				case clientv3.EventTypeDelete:
					fmt.Printf("[键删除] %s\n", string(event.Kv.Key))
				}
			}
		}
	}
}

// PrintAllData 打印所有数据（树状结构）
func (em *EtcdManager) PrintAllData(ctx context.Context) error {
	fmt.Println("=== ETCD 所有数据 ===")

	// 获取所有键值对
	allData, err := em.GetAllKeys(ctx, "")
	if err != nil {
		return err
	}
	if len(allData) == 0 {
		fmt.Println("ETCD没有数据")
		return nil
	}

	// 构建树状结构
	tree := make(map[string]interface{})

	for key, value := range allData {
		parts := strings.Split(key, "/")
		current := tree

		for i, part := range parts {
			if part == "" {
				continue
			}

			if i == len(parts)-1 {
				// 叶子节点，存储值
				current[part] = value
			} else {
				// 中间节点
				if current[part] == nil {
					current[part] = make(map[string]interface{})
				}
				current = current[part].(map[string]interface{})
			}
		}
	}

	// 打印树状结构
	em.printTree(tree, 0)
	fmt.Printf("总计: %d 个键值对\n", len(allData))

	return nil
}

// printTree 递归打印树状结构
func (em *EtcdManager) printTree(node map[string]interface{}, depth int) {
	prefix := strings.Repeat("  ", depth)

	for key, value := range node {
		switch v := value.(type) {
		case string:
			fmt.Printf("%s├── %s: %s\n", prefix, key, v)
		case map[string]interface{}:
			fmt.Printf("%s├── %s/\n", prefix, key)
			em.printTree(v, depth+1)
		}
	}
}

// PutKey 写入键值对（用于测试）
func (em *EtcdManager) PutKey(ctx context.Context, key, value string) error {
	_, err := em.client.Put(ctx, key, value)
	return err
}

// DeleteKey 删除键（用于测试）
func (em *EtcdManager) DeleteKey(ctx context.Context, key string) error {
	_, err := em.client.Delete(ctx, key)
	return err
}

func main() {
	// etcd 连接端点
	endpoints := []string{"104.168.125.34:2379"}

	// 创建 etcd 管理器
	etcdMgr, err := NewEtcdManager(endpoints)
	if err != nil {
		log.Fatalf("创建 etcd 管理器失败: %v", err)
	}
	defer etcdMgr.Close()

	ctx := context.Background()

	// 1. 获取并打印所有数据
	fmt.Println("1. 获取所有 etcd 数据:")
	err = etcdMgr.PrintAllData(ctx)
	if err != nil {
		log.Printf("获取所有数据失败: %v", err)
	}

	// 2. 列出所有键
	fmt.Println("\n2. 列出所有键:")
	keys, err := etcdMgr.ListKeys(ctx, "")
	if err != nil {
		log.Printf("列出键失败: %v", err)
	} else {
		fmt.Printf("找到 %d 个键:\n", len(keys))
		for i, key := range keys {
			fmt.Printf("  %d. %s\n", i+1, key)
		}
	}

	// 3. 获取指定前缀的所有键值对（示例）
	fmt.Println("\n3. 获取指定前缀的键值对:")
	prefixData, err := etcdMgr.GetAllKeys(ctx, "/registry/pods") // 修改为你需要的前缀
	if err != nil {
		log.Printf("获取前缀数据失败: %v", err)
	} else {
		fmt.Printf("找到 %d 个键值对:\n", len(prefixData))
		for key, value := range prefixData {
			// 只显示前100个字符，避免输出过长
			if len(value) > 100 {
				value = value[:100] + "..."
			}
			fmt.Printf("  %s: %s\n", key, value)
		}
	}

	// 4. 启动多个监视器
	fmt.Println("\n4. 启动监视器...")

	// 创建取消上下文，用于优雅停止
	watchCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 监视所有键的变化
	go etcdMgr.WatchPrefix(watchCtx, "")

	// 监视特定前缀（示例）
	go etcdMgr.WatchPrefix(watchCtx, "/registry/")

	// 监视特定键（示例）
	// go etcdMgr.WatchSpecificKey(watchCtx, "/registry/config")

	// 5. 测试：写入一些数据观察 watch 效果
	fmt.Println("\n5. 开始测试数据变更（10秒后开始）...")
	time.Sleep(10 * time.Second)

	testKey := "/test/watch/demo"

	// 测试 PUT 操作
	fmt.Println(">>> 测试 PUT 操作")
	err = etcdMgr.PutKey(ctx, testKey, "初始值")
	if err != nil {
		log.Printf("PUT 失败: %v", err)
	}
	time.Sleep(2 * time.Second)

	// 测试 UPDATE 操作
	fmt.Println(">>> 测试 UPDATE 操作")
	err = etcdMgr.PutKey(ctx, testKey, "更新后的值")
	if err != nil {
		log.Printf("UPDATE 失败: %v", err)
	}
	time.Sleep(2 * time.Second)

	// 测试 DELETE 操作
	fmt.Println(">>> 测试 DELETE 操作")
	err = etcdMgr.DeleteKey(ctx, testKey)
	if err != nil {
		log.Printf("DELETE 失败: %v", err)
	}

	// 6. 持续运行一段时间观察变化
	fmt.Println("\n6. 持续监视中（30秒后自动停止）...")
	time.Sleep(30 * time.Second)

	cancel() // 停止所有监视器
	fmt.Println("程序结束")
}
