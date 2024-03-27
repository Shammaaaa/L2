package main

import (
	"bufio"
	"context"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const (
	timeoutPrefix  = "--timeout="
	defaultTimeout = 10 * time.Second
)

var (
	errNotEnoughArguments = errors.New("not enough arguments")
	errInvalidArgumentKey = errors.New("invalid argument key ")
	errInvalidTimeout     = errors.New("invalid timeout")

	errClosedByServer = errors.New("closed by server")
	errClosedByUs     = errors.New("closed by us")
)

type args struct {
	timeout time.Duration
	host    string
	port    string
}

func main() {
	args, err := parseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), args.timeout)
	defer cancel()

	err = runTelnet(ctx, os.Stdout, os.Stdin, args.host, args.port)
	if err != nil && err != errClosedByUs {
		log.Fatal(err)
	}
}

func parseArgs(rawArgs []string) (args, error) {
	timeout := defaultTimeout

	if len(rawArgs) == 3 {
		if !strings.HasPrefix(rawArgs[0], timeoutPrefix) {
			return args{}, errInvalidArgumentKey
		}

		dStr := rawArgs[0][len(timeoutPrefix):]
		var err error
		timeout, err = time.ParseDuration(dStr)
		if err != nil {
			return args{}, errInvalidTimeout
		}
		rawArgs = rawArgs[1:]
	}

	if len(rawArgs) != 2 {
		return args{}, errNotEnoughArguments
	}

	host := rawArgs[0]
	port := rawArgs[1]
	return args{timeout, host, port}, nil
}

func runTelnet(ctx context.Context, out io.Writer, in io.Reader, host, port string) error {
	con, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return err
	}

	readerChan := bufferedReaderToChan(ctx, in)
	serverChan := bufferedReaderToChan(ctx, con)

	for err == nil {

		select {
		case <-ctx.Done():
			err = ctx.Err()
		case data, ok := <-readerChan:
			if ok {
				_, err = con.Write(data)
			} else {
				err = errClosedByUs
			}
		case data, ok := <-serverChan:
			if ok {
				_, err = out.Write(data)
			} else {
				err = errClosedByServer
			}
		}
	}
	return err
}

func bufferedReaderToChan(ctx context.Context, r io.Reader) <-chan []byte {
	ch := make(chan []byte)

	go func() {
		sc := bufio.NewScanner(r)
		for sc.Scan() && sc.Err() == nil && ctx.Err() == nil {
			data := sc.Text()
			if len(data) == 0 {
				break
			}
			ch <- []byte(data + "\n")
		}
		close(ch)
	}()

	return ch
}
