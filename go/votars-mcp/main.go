package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/scarletlabs-ai/Votars-MCP/go/votars-mcp/mcp"
)

func main() {
	var (
		transport, port string
	)
	flag.StringVar(&transport, "t", "sse", "Transport type (stdio or sse)")
	flag.StringVar(&port, "p", "8080", "Listen on port (only for `sse`)")
	flag.Parse()

	s := mcp.NewMCPServer()

	switch transport {
	case "stdio":
		if err := s.ServeStdio(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	case "sse":
		addr := fmt.Sprintf("0.0.0.0:%s", port)
		log.Println("SSE server listening on", addr)
		if err := s.ServeSSE(addr); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	default:
		log.Fatalf(
			"Invalid transport type: %s. Must be 'stdio' or 'sse'",
			transport,
		)
	}
}
