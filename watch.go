package errors

import (
	"github.com/go-mysql-org/go-mysql/canal"
	siddontang_log "github.com/siddontang/go-log/log"
	"os"
	"regexp"
)

var cc *canal.Canal

func initCanal(dsn string) {
	params := regexp.MustCompile(
		`^(?:(?P<user>.*?)(?::(?P<passwd>.*))?@)?` +
			`(?:(?P<net>[^\(]*)(?:\((?P<addr>[^\)]*)\))?)?` +
			`\/(?P<dbname>.*?)` +
			`(?:\?(?P<params>[^\?]*))?$`).FindStringSubmatch(dsn)
	cfg := canal.NewDefaultConfig()
	cfg.User = params[1]
	cfg.Password = params[2]
	cfg.Addr = params[4]
	cfg.Dump.ExecutionPath = ""

	h, _ := siddontang_log.NewStreamHandler(os.Stdout)
	l := siddontang_log.NewDefault(h)
	l.SetLevel(siddontang_log.LevelWarn)
	cfg.Logger = l
	c, err := canal.NewCanal(cfg)
	if err != nil {
		zapLogger.Sugar().Panicf("%+v", err)
	}
	c.SetEventHandler(&handler{
		dsn: dsn,
	})

	cc = c
}

func watch() {
	pos, err := cc.GetMasterPos()
	if err != nil {
		zapLogger.Sugar().Panicf("%+v", err)
	}

	_ = cc.RunFrom(pos)
}

func closeWatch() {
	cc.Close()
}

type handler struct {
	dsn string
	canal.DummyEventHandler
}

func (s *handler) OnRow(e *canal.RowsEvent) error {
	params := regexp.MustCompile(
		`^(?:(?P<user>.*?)(?::(?P<passwd>.*))?@)?` +
			`(?:(?P<net>[^\(]*)(?:\((?P<addr>[^\)]*)\))?)?` +
			`\/(?P<dbname>.*?)` +
			`(?:\?(?P<params>[^\?]*))?$`).FindStringSubmatch(s.dsn)
	if e.Table.Schema == params[5] && e.Table.Name == "error_codes" && (e.Action == "insert" || e.Action == "update" || e.Action == "delete") {
		register()
	}
	return nil
}

func (*handler) String() string {
	return "handler"
}
