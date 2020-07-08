package immuadmin

import (
	"bytes"
	"context"
	"github.com/codenotary/immudb/pkg/api/schema"
	"github.com/codenotary/immudb/pkg/server"
	"github.com/codenotary/immudb/pkg/server/servertest"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"io/ioutil"
	"testing"
)

func TestCommandLine_PrintTreeInit(t *testing.T) {
	bs := servertest.NewBufconnServer(server.Options{}.WithAuth(false).WithInMemoryStore(true))
	bs.Start()
	dialOptions := []grpc.DialOption{
		grpc.WithContextDialer(bs.Dialer), grpc.WithInsecure(),
	}
	opts := Options()
	opts.DialOptions = &dialOptions
	cmdl := commandline{
		context: context.Background(),
		options: opts,
	}
	_ = cmdl.connect(&cobra.Command{}, []string{})

	cmdl.printTree(&cobra.Command{})

}

func TestCommandLine_PrintTree(t *testing.T) {
	bs := servertest.NewBufconnServer(server.Options{}.WithAuth(false).WithInMemoryStore(true))
	bs.Start()
	bs.Server.Set(context.TODO(), &schema.KeyValue{Key: []byte(`myFirstElementKey`), Value: []byte(`firstValue`)})

	cmd := cobra.Command{}
	cl := new(commandline)
	cl.context = context.Background()
	dialOptions := []grpc.DialOption{
		grpc.WithContextDialer(bs.Dialer), grpc.WithInsecure(),
	}

	cl.options = Options()
	//cl.options.WithAuth(false)
	cl.options.DialOptions = &dialOptions
	cl.printTree(&cmd)

	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"print"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Contains(t, string(out), "Immudb Merkle Tree")

}
