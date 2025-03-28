![Votars Logo](https://votars.ai/_next/static/media/logo.e7b6bff6.svg) 
# Votars MCP 
[![smithery badge](https://smithery.ai/badge/@scarletlabs-ai/votars-mcp)](https://smithery.ai/server/@scarletlabs-ai/votars-mcp)


## Overview

Votars-MCP is a tool that supports multiple language implementations of the **Votars MCP server**. Currently, only the Go version is available, with other languages to be added in future releases. It supports two interaction modes: `sse` (Server-Sent Events) and `stdio` (Standard Input/Output). It is designed to provide seamless integration with the Votars AI platform for processing various tasks.

## About Votars

[Votars](https://votars.ai/en/) is the world's smartest multilingual meeting assistant, designed for voice recording, transcription, and advanced AI processing. It features real-time translation, intelligent error correction, AI summarization, smart content generation, and AI discussions. The Votars app is available on [Web](https://votars.ai/en/), [iOS](https://apps.apple.com/us/app/votars-ai-transcribe-organize/id6737496290), and [Android](https://play.google.com/store/apps/details?id=com.votars.transcribe).

Additionally, Votars is an AI-powered platform that enables developers to integrate advanced AI functionalities into their applications. By leveraging Votars, you can process complex tasks efficiently with robust APIs designed for high performance and scalability.

## Features
- **Easy Integration with Votars**
- **Modular Design:** Ready to be extended with additional functionalities.
- **Supported MCP Tools:**
  - `Votars_fetch_recent_transcripts`: Allows users to read recent transcripts from their workspace, providing convenient access to the latest recorded sessions.
  - `Votars_fetch_a_specific_transcript`: Enables users to retrieve specific transcripts by providing a transcript ID, allowing targeted retrieval of stored data.
  
  More functionalities will be added soon. Stay tuned!

## Installation (Go MCP)

### Installing via Smithery

To install votars-mcp for Claude Desktop automatically via [Smithery](https://smithery.ai/server/@scarletlabs-ai/votars-mcp):

```bash
npx -y @smithery/cli install @scarletlabs-ai/votars-mcp --client claude
```

### Manual Installation
To install the Go version of Votars MCP from the GitHub repository, use:

```bash
 go install github.com/scarletlabs-ai/Votars-MCP/go/votars-mcp@latest
```

## Usage (Go MCP)

### Run MCP Service
Before using the `sse` mode, you need to run the MCP server. Open a terminal and run:

```bash
votars-mcp -t sse -p 8080
```

This command starts the MCP service on port 8080, ready to accept `sse` requests.


### 1. SSE Mode

For `sse` mode, you need to provide the API key via request headers in the configuration file.

Configuration file example (`mcp.config.json`):

```json
{
  "mcpServers": {
    "Votars MCP": {
      "type": "sse",
      "url": "http://0.0.0.0:8080/sse",
      "headers": {
        "Authorization": "Bearer <your-api-key>"
      }
    }
  }
}
```

### 2. Stdio Mode

For `stdio` mode, set the API key as an environment variable.


Configuration file example (`mcp.config.json`):

```json
{
  "mcpServers": {
    "Votars MCP Stdio": {
      "type": "stdio",
      "command": "votars-mcp",
      "args": ["-t", "stdio"],
      "env": {
        "VOTARS_API_KEY": "<your-api-key>"
      }
    }
  }
}
```

## Obtaining Your API Key

1. Go to [Votars.AI](https://votars.ai/en/) and register.
2. Navigate to your workspace's `Settings`.
3. Create an API Key under the API Key management section.

![manage apikey](https://private-user-images.githubusercontent.com/677477/427500562-8cfd8465-f408-4e9b-a101-8b9b8e5e57f5.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3NDMwNzQ4NzAsIm5iZiI6MTc0MzA3NDU3MCwicGF0aCI6Ii82Nzc0NzcvNDI3NTAwNTYyLThjZmQ4NDY1LWY0MDgtNGU5Yi1hMTAxLThiOWI4ZTVlNTdmNS5wbmc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjUwMzI3JTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI1MDMyN1QxMTIyNTBaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT02ODkyMWY3YzgyYTA2ZjdmNGQxN2MyYzllMzBmZmFiZmVjNGFmZTliNDQzODUwMjU2M2E1MjJkZTI4MmExM2VmJlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCJ9.V0GAoh6l0JFRxmSokliGsVt5yNpxtmTmeCFiZG7U3jU)

## Roadmap

- **Current Support:** Go
- **Planned Support:** Python, JavaScript, Rust, etc.

## License

MIT License
