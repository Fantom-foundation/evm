package proxy

import (
	"time"

	"github.com/andrecronje/evm/service"
	"github.com/andrecronje/evm/state"
	proxy "github.com/andrecronje/lachesis/src/proxy/lachesis"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

//------------------------------------------------------------------------------

type Config struct {
	proxyAddr    string //bind address of this app proxy
	lachesisAddr string //address of node
	apiAddr      string //address of HTTP API service
	ethDir       string //directory containing eth config
	pwdFile      string //file containing password to unlock ethereum accounts
	databaseFile string //file containing LevelDB database
	cache        int    //Megabytes of memory allocated to internal caching (min 16MB / database forced)
	timeout      time.Duration
}

func NewConfig(proxyAddr,
	lachesisAddr,
	apiAddr,
	ethDir,
	pwdFile,
	dbFile string,
	cache int,
	timeout time.Duration) Config {

	return Config{
		proxyAddr:    proxyAddr,
		lachesisAddr: lachesisAddr,
		apiAddr:      apiAddr,
		ethDir:       ethDir,
		pwdFile:      pwdFile,
		databaseFile: dbFile,
		cache:        cache,
		timeout:      timeout,
	}
}

//------------------------------------------------------------------------------

type Proxy struct {
	service       *service.Service
	state         *state.State
	lachesisProxy *proxy.WebsocketLachesisProxy
	submitCh      chan []byte
	logger        *logrus.Logger
}

func NewProxy(config Config, logger *logrus.Logger) (*Proxy, error) {
	submitCh := make(chan []byte)

	logger.Debug("state.NewState")
	state_, err := state.NewState(logger, config.databaseFile, config.cache)
	if err != nil {
		log.WithError(err).Error("error building state")
		return nil, err
	}

	logger.Debug("service.NewService")
	service_ := service.NewService(config.ethDir,
		config.apiAddr,
		config.pwdFile,
		state_,
		submitCh,
		logger)

	logger.Debug("proxy.NewSocketLachesisProxy")
	lachesisProxy, err := proxy.NewWebsocketLachesisProxy(config.lachesisAddr,
		config.timeout,
		logger)
	if err != nil {
		log.WithError(err).Error("error building socket proxy")
		return nil, err
	}

	logger.Debug("Return &Proxy")
	return &Proxy{
		service:       service_,
		state:         state_,
		lachesisProxy: lachesisProxy,
		submitCh:      submitCh,
		logger:        logger,
	}, nil
}

func (p *Proxy) Run() error {

	go p.service.Run()

	p.Serve()

	return nil
}

func (p *Proxy) Serve() {
	for {
		select {
		case tx := <-p.submitCh:
			p.logger.Debug("Proxy about to submit tx")
			if err := p.lachesisProxy.SubmitTx(tx); err != nil {
				p.logger.WithError(err).Error("SubmitTx")
			}
			p.logger.Debug("Proxy submitted tx")
		case commit := <-p.lachesisProxy.CommitCh():
			p.logger.Debug("CommitBlock")
			stateHash, err := p.state.ProcessBlock(commit.Block)
			commit.Respond(stateHash.Bytes(), err)
		}
	}
}
