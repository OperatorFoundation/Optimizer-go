package optimizer

import (
	"github.com/OperatorFoundation/Shadow-go/shadow/v3"
	"github.com/kataras/golog"
	"net"
	"os"
	"testing"
)

const data = "test"

func TestMain(m *testing.M) {
	config := shadow.NewServerConfig("1234", "DarkStar")
	listener, listenErr := config.Listen("127.0.0.1:1235")
	if listenErr != nil {
		return
	}
	go acceptConnections(listener)

	os.Exit(m.Run())
}

func acceptConnections(listener net.Listener) {
	for {
		_, err := listener.Accept()
		if err != nil {
			return
		}
	}
}

func TestShadowDial(t *testing.T) {
	MakeLog()
	shadowTransport := shadow.Transport{Password: "1234", CipherName: "CHACHA20-IETF-POLY1305", Address: "127.0.0.1:1235"}
	_, err := shadowTransport.Dial()
	if err != nil {
		t.Fail()
	}
}

func TestOptimizerShadowDial(t *testing.T) {
	MakeLog()
	shadowTransport := shadow.NewTransport("1234", "DarkStar", "127.0.0.1:1235")
	transports := []TransportDialer{&shadowTransport}
	strategy := NewFirstStrategy(transports)
	optimizerTransport := NewOptimizerClient(transports, strategy)
	_, err := optimizerTransport.Dial()
	if err != nil {
		t.Fail()
	}
}

func TestOptimizerTransportFirstDial(t *testing.T) {
	MakeLog()
	shadowTransport := shadow.NewTransport("1234", "CHACHA20-IETF-POLY1305", "127.0.0.1:1235")
	transports := []TransportDialer{&shadowTransport}
	strategy := NewFirstStrategy(transports)
	optimizerTransport := NewOptimizerClient(transports, strategy)
	for i := 1; i <= 3; i++ {
		_, err := optimizerTransport.Dial()
		if err != nil {
			t.Fail()
		}
	}
}

func TestOptimizerTransportRandomDial(t *testing.T) {
	MakeLog()
	shadowTransport := shadow.NewTransport("1234", "CHACHA20-IETF-POLY1305", "127.0.0.1:1235")
	transports := []TransportDialer{&shadowTransport}
	strategy := NewRandomStrategy(transports)
	optimizerTransport := NewOptimizerClient(transports, strategy)

	for i := 1; i <= 3; i++ {
		_, err := optimizerTransport.Dial()
		if err != nil {
			t.Fail()
		}
	}
}

func TestOptimizerTransportRotateDial(t *testing.T) {
	MakeLog()
	shadowTransport := shadow.NewTransport("1234", "CHACHA20-IETF-POLY1305", "127.0.0.1:1235")
	transports := []TransportDialer{&shadowTransport}
	strategy := NewRotateStrategy(transports)
	optimizerTransport := NewOptimizerClient(transports, strategy)

	for i := 1; i <= 3; i++ {
		_, err := optimizerTransport.Dial()
		if err != nil {
			t.Fail()
		}
	}
}

func TestOptimizerTransportTrackDial(t *testing.T) {
	MakeLog()
	shadowTransport := shadow.NewTransport("1234", "CHACHA20-IETF-POLY1305", "127.0.0.1:1235")
	transports := []TransportDialer{&shadowTransport}
	strategy := NewTrackStrategy(transports)
	optimizerTransport := NewOptimizerClient(transports, strategy)

	for i := 1; i <= 3; i++ {
		_, err := optimizerTransport.Dial()
		if err != nil {
			t.Fail()
		}
	}
}

func TestOptimizerTransportMinimizeDialDurationDial(t *testing.T) {
	MakeLog()
	shadowTransport := shadow.NewTransport("1234", "CHACHA20-IETF-POLY1305", "127.0.0.1:1235")
	transports := []TransportDialer{&shadowTransport}
	strategy := NewMinimizeDialDuration(transports)
	optimizerTransport := NewOptimizerClient(transports, strategy)

	for i := 1; i <= 3; i++ {
		_, err := optimizerTransport.Dial()
		if err != nil {
			t.Fail()
		}
	}
}

func MakeLog() {
	golog.SetLevel("debug")
	golog.SetOutput(os.Stderr)
}
