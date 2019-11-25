package transport

import (
	"testing"
	"github.com/micro/go-micro/transport"
)

func TestMeTransport(t *testing.T)  {
	tr := transport.NewTransport()
	//bind / listen
	l, err := tr.Listen("127.0.0.1:8889")
	if err != nil{
		t.Fatalf("Unexpected error listening %v", err)
	}
	defer l.Close()

	//accept
	go func() {
		if err := l.Accept(func(socket transport.Socket) {
			for{
					var m transport.Message
					if err := socket.Recv(&m); err != nil{
						return
					}
					t.Logf("Server Received %s", string(m.Body))
					if err := socket.Send(&transport.Message{
						Header: nil,
						Body:   []byte("ping"),
					}); err != nil{
						return
					}
			}
		});err != nil{
			t.Fatalf("Unexpected error accepting %v", err)
		}
	}()

	//Dial
	c, err := tr.Dial("127.0.0.1:8888")
	if err != nil {
		t.Fatalf("Unexpected error dialing %v", err)
	}
	defer c.Close()

	// send <=> receive
	for i := 0; i < 3; i++ {
		if err := c.Send(&transport.Message{
			Body: []byte(`ping`),
		}); err != nil {
			return
		}
		var m transport.Message
		if err := c.Recv(&m); err != nil {
			return
		}
		t.Logf("Client Received %s", string(m.Body))
	}
}