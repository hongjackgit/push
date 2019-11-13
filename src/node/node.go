package push



// Node 执行 cron 命令服务的结构体
type Service struct {
	ttl  int64
	lID  client.LeaseID // lease id
	done chan struct{}
	vids map[string]int64  //key:task/cms/platform,val:1,2,3
}