package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func handler(in []byte) []byte {
	return in
}

func listenSocket(socketName string) error {
	l, err := net.Listen("unix", socketName)
	if err != nil {
		return err
	}
	defer l.Close()

	log.Printf("listen on socket %s", socketName)

	connCh := make(chan net.Conn)
	errCh := make(chan error)
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				errCh <- err
				return
			}
			connCh <- conn
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case conn := <-connCh:
			log.Print("connected")
			go func() {
				defer conn.Close()
				buf := make([]byte, 512)
				for {
					if _, err := conn.Read(buf); err != nil {
						log.Print(err)
						return
					}
					if _, err := conn.Write(handler(buf)); err != nil {
						log.Print(err)
						return
					}
				}
			}()
		case err := <-errCh:
			return err
		case <-sigs:
			return nil
		}
	}
}

var rootCmd = &cobra.Command{
	Use: "start-process-vs-unix-domain-socket",
	RunE: func(cmd *cobra.Command, args []string) error {
		daemon, err := cmd.Flags().GetBool("server")
		if err != nil {
			return err
		}
		if daemon {
			socketName, err := cmd.Flags().GetString("socket")
			if err != nil {
				return err
			}
			if err := listenSocket(socketName); err != nil {
				return err
			}
		} else {
			value, err := cmd.Flags().GetString("value")
			if err != nil {
				return err
			}
			fmt.Println(handler([]byte(value)))
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Bool("server", false, "Wait for the value with socket")
	rootCmd.Flags().String("socket", "rpos.sock", "Socket name to be listen by the server")
	rootCmd.Flags().String("value", "", "Value to be processed")
}
