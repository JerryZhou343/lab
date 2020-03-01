package di

import (
	"context"
	"errors"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/naming"
	"github.com/bilibili/kratos/pkg/naming/discovery"
	"github.com/bilibili/kratos/pkg/naming/etcd"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/mfslog/kratos-ddd/api/rpc"
	"github.com/mfslog/kratos-ddd/infra/conf"
	"go.etcd.io/etcd/clientv3"
	"net"
	"os"
	"strings"
	"time"
)

var (
	nsEtcd      *etcd.EtcdBuilder
	nsDiscovery *discovery.Discovery
)

type App struct {
	svcHandler *rpc.Handler
	grpc       *warden.Server
	config     *conf.Config
	ds         *etcd.EtcdBuilder
}

func NewApp(svcHandler *rpc.Handler, g *warden.Server, config *conf.Config) (app *App, closeFunc func(), err error) {
	var (
		unregister context.CancelFunc
		ip         string
	)
	app = &App{
		svcHandler: svcHandler,
		grpc:       g,
		config:     config,
	}

	//服务注册
	hostName, _ := os.Hostname()
	ip, err = app.getLocalIP()
	if err != nil {
		return
	}
	addrInfo := strings.Split(config.Grpc.Addr, ":")
	if len(addrInfo) != 2 {
		err = errors.New("server addr config error :" + config.Grpc.Addr)
		return
	}
	addrs := []string{fmt.Sprintf("%s:%s", ip, addrInfo[1])}
	ins := &naming.Instance{
		Region:   env.Region,
		Zone:     env.Zone,
		Env:      env.DeployEnv,
		AppID:    env.AppID,
		Hostname: hostName,
		Addrs:    addrs,
		Version:  "",
		LastTs:   0,
		Metadata: nil,
		Status:   0,
	}

	//unregister, err = app.registerDiscovery(ins)
	unregister, err = app.registerEtcd(ins)
	closeFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		if err := g.Shutdown(ctx); err != nil {
			log.Error("grpcSrv.Shutdown error(%v)", err)
		}
		cancel()
		unregister()
	}
	return
}

func (a *App) registerDiscovery(ins *naming.Instance) (unregister context.CancelFunc, err error) {
	nsDiscovery = discovery.New(nil)
	unregister, err = nsDiscovery.Register(context.Background(), ins)
	return
}

func (a *App) registerEtcd(ins *naming.Instance) (unregister context.CancelFunc, err error) {
	nodes := strings.Split(env.DiscoveryNodes, ",")
	cfg := &clientv3.Config{
		Endpoints:            nodes,
		AutoSyncInterval:     time.Second,
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    time.Second,
		DialKeepAliveTimeout: 6 * time.Second,
		Context:              context.Background(),
	}

	nsEtcd, err = etcd.New(cfg)
	unregister, err = nsEtcd.Register(context.Background(), ins)
	return
}

func (a *App) getLocalIP() (addr string, err error) {
	var (
		netInterfaces []net.Interface
	)
	netInterfaces, err = net.Interfaces()
	if err != nil {
		return
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						addr = ipnet.IP.String()
						return
					}
				}
			}
		}
	}
	return
}
