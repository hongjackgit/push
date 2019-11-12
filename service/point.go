package service

type Point struct {
	ID       string `json:"id"`  // machine id
	PID      string `json:"pid"` // 进程 pid
	PIDFile  string `json:"-"`
	IP       string `json:"ip"` // node ip
	Hostname string `json:"hostname"`

	UpTime   time.Time `json:"up"`     // 启动时间
	DownTime time.Time `json:"down"` // 上次关闭时间

	Alived    bool `json:"alived"` // 是否可用
	Connected bool `json:"connected"`   // 当 Alived 为 true 时有效，表示心跳是否正常
}

func (p *Point) Start() (*client.DeleteResponse, error) {
	
}

func (p *Point) Put(opts ...client.OpOption) (*client.PutResponse, error) {
}

func (p *Point) Del() (*client.DeleteResponse, error) {
}

// 判断 node 是否已注册到 etcd
// 存在则返回进行 pid，不存在返回 -1
func (p *Point) Exist() (pid int, err error) {
	resp, err := DefalutClient.Get(conf.Config.Node + n.ID)
	if err != nil {
		return
	}

	if len(resp.Kvs) == 0 {
		return -1, nil
	}

	if pid, err = strconv.Atoi(string(resp.Kvs[0].Value)); err != nil {
		if _, err = DefalutClient.Delete(conf.Config.Node + n.ID); err != nil {
			return
		}
		return -1, nil
	}

	p, err := os.FindProcess(pid)
	if err != nil {
		return -1, nil
	}

	// TODO: 暂时不考虑 linux/unix 以外的系统
	if p != nil && p.Signal(syscall.Signal(0)) == nil {
		return
	}

	return -1, nil
}