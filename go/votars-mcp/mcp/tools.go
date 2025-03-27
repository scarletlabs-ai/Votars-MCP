package mcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
)

func makeRequest(ctx context.Context, method, endpoint, token string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", APIDomain, endpoint)
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("VOTARS-API-KEY", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

const (
	APIDomain                        = "https://api.votars.ai"
	EndpointGetRecentTranscripts     = "/v1/mcp/tools/transcripts/recent/"
	EndpointGetTranscriptByDocIDTmpl = "/v1/mcp/tools/transcripts/%d/"
)

func FetchRecentTranscripts(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	token, err := tokenFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("missing token: %v", err)
	}
	data, err := makeRequest(ctx, http.MethodGet, EndpointGetRecentTranscripts, token)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(string(data)), nil
}

type transcriptResp struct {
	Content string `json:"content"`
}

func FetchSpecificTranscript(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	docIDNumber, ok := request.Params.Arguments["id"].(float64)
	if !ok {
		return nil, errors.New("id must be a number")
	}

	docID := int64(docIDNumber)
	token, err := tokenFromContext(ctx)
	if err != nil {
		return nil, err
	}
	data, err := makeRequest(ctx, http.MethodGet, fmt.Sprintf(EndpointGetTranscriptByDocIDTmpl, docID), token)
	if err != nil {
		return nil, err
	}
	transcript := &transcriptResp{}
	if err := json.Unmarshal(data, transcript); err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(transcript.Content), nil
}
