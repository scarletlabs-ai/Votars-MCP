package mcp

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type MCPServer struct {
	server *server.MCPServer
}

func NewMCPServer() *MCPServer {
	s := server.NewMCPServer(
		"Votars MCP",
		"0.9.0",
		server.WithToolCapabilities(true),
	)
	fetchRecentTranscriptsTool := mcp.NewTool("Votars fetch recent transcripts",
		mcp.WithDescription("Retrive recent transcripts from workspace"),
	)
	fetchSpecificTranscriptTool := mcp.NewTool("Votars fetch a specific transcript",
		mcp.WithDescription("Retrive the transcript from the workspace by its ID"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Transcript ID to retrieve"),
		),
	)
	s.AddTool(fetchRecentTranscriptsTool, FetchRecentTranscripts)
	s.AddTool(fetchSpecificTranscriptTool, FetchSpecificTranscript)
	return &MCPServer{
		server: s,
	}
}

func (s *MCPServer) ServeSSE(addr string) error {
	return server.NewSSEServer(s.server,
		server.WithBaseURL(fmt.Sprintf("http://%s", addr)),
		server.WithSSEContextFunc(authFromRequest),
	).Start(addr)
}

func (s *MCPServer) ServeStdio() error {
	return server.ServeStdio(s.server, server.WithStdioContextFunc(authFromEnv))
}

func authFromEnv(ctx context.Context) context.Context {
	return withAuthKey(ctx, os.Getenv("VOTARS_API_KEY"))
}

type authKey struct{}

func withAuthKey(ctx context.Context, auth string) context.Context {
	return context.WithValue(ctx, authKey{}, auth)
}

func authFromRequest(ctx context.Context, r *http.Request) context.Context {
	authHeader := r.Header.Get("Authorization")
	token := strings.Replace(authHeader, "Bearer ", "", 1)
	return withAuthKey(ctx, token)
}

func tokenFromContext(ctx context.Context) (string, error) {
	auth, ok := ctx.Value(authKey{}).(string)
	if !ok {
		return "", fmt.Errorf("missing auth")
	}
	return auth, nil
}
