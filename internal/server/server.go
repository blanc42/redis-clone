package server

import (
	"bufio"
	"fmt"
	"net"
	"redis-clone/internal/commands"
	"redis-clone/internal/datastore"
	"redis-clone/pkg/protocol"
)

type Server struct {
	addr     string
	store    datastore.DataStore
	commands map[string]commands.Commander
}

func NewServer(addr string) *Server {
	store := datastore.NewInMemoryStore()
	return &Server{
		addr:     addr,
		store:    store,
		commands: makeCommands(store),
	}
}

func makeCommands(store datastore.DataStore) map[string]commands.Commander {
	base := commands.NewBaseCommand(store)
	return map[string]commands.Commander{
		"GET": &commands.GetCommand{BaseCommand: base},
		"SET": &commands.SetCommand{BaseCommand: base},
		// Add more commands here
	}
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		cmd, err := protocol.DecodeRESP(reader)
		if err != nil {
			fmt.Println("Error decoding RESP:", err)
			return
		}

		if len(cmd) == 0 {
			continue
		}

		cmdName := cmd[0]
		cmdArgs := cmd[1:]

		command, ok := s.commands[cmdName]
		if !ok {
			protocol.EncodeRESP(conn, fmt.Sprintf("Unknown command: %s", cmdName))
			continue
		}

		result, err := command.Execute(cmdArgs)
		if err != nil {
			protocol.EncodeRESP(conn, fmt.Sprintf("Error: %s", err))
		} else {
			protocol.EncodeRESP(conn, result)
		}
	}
}
