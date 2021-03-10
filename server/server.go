package server

import "fmt"

var (
	WSSHandler      = func(bot int64, data []byte) []byte { fmt.Println(string(data)); return []byte("ok") }
	WSCHandler      = func(bot int64, data []byte) []byte { fmt.Println(string(data)); return []byte("ok") }
	HttpPostHandler = func(bot int64, send []byte, data []byte) { fmt.Println(string(data)) }
	HttpHandler     = func(bot int64, path string, data []byte) []byte { fmt.Println(string(data)); return []byte("ok") }

	CoreInfo  = func(s string, v ...interface{}) { fmt.Printf(s, v...) }
	CoreDebug = func(s string, v ...interface{}) { fmt.Printf(s, v...) }
)

func (s *WSC) INFO(text interface{}) {
	CoreInfo("[INFO][WSC] INFO: %v ID: %d URL: %s ", text, s.ID, s.Addr)
}

func (s *WSC) DEBUG(text interface{}) {
	CoreDebug("[DEBUG][WSC] ID: %d URL: %s DEBUG: %v", s.ID, s.Addr, text)
}

func (s *WSC) ERROR(text interface{}) {
	CoreInfo("[ERROR][WSC] ERROR: %v ID: %d URL: %s ", text, s.ID, s.Addr)
}

func (s *WSC) PANIC(err interface{}, traceback []byte) {
	CoreInfo("[PANIC][WSC] ID: %d URL: %s \n[ERROR]\n%v\n[TRACEBACK]\n%s", s.ID, s.Addr, err, string(traceback))
}

func (s *WSS) INFO(text interface{}) {
	CoreInfo("[INFO][WSS] INFO: %v ID: %d URL: %s ", text, s.ID, s.Addr)
}

func (s *WSS) DEBUG(text interface{}) {
	CoreDebug("[DEBUG][WSS] ID: %d URL: %s DEBUG: %v", s.ID, s.Addr, text)
}

func (s *WSS) ERROR(text interface{}) {
	CoreInfo("[ERROR][WSS] ERROR: %v ID: %d URL: %s ", text, s.ID, s.Addr)
}

func (s *WSS) PANIC(err interface{}, traceback []byte) {
	CoreInfo("[PANIC][WSS] ID: %d URL: %s \n[ERROR]\n%v\n[TRACEBACK]\n%s", s.ID, s.Addr, err, string(traceback))
}

func (s *HTTP) INFO(text interface{}) {
	CoreInfo("[INFO][HTTP] INFO: %v ID: %d URL: %s ", text, s.ID, s.Addr)
}

func (s *HTTP) DEBUG(text interface{}) {
	CoreDebug("[DEBUG][HTTP] ID: %d URL: %s DEBUG: %v", s.ID, s.Addr, text)
}

func (s *HTTP) ERROR(text interface{}) {
	CoreInfo("[ERROR][HTTP] ERROR: %v ID: %d URL: %s ", text, s.ID, s.Addr)
}

func (s *HTTP) PANIC(err interface{}, traceback []byte) {
	CoreInfo("[PANIC][HTTP] ID: %d URL: %s \n[ERROR]\n%v\n[TRACEBACK]\n%s", s.ID, s.Addr, err, string(traceback))
}

type Server interface {
	Run()
	Close()
	Send(data []byte)
}