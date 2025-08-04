# Log-Aggregator-status-go

log-aggregator/
├── cmd/                     # Entrypoint(s)
│   └── server/              # Main service
│       └── main.go
├── internal/                # Private application code
│   ├── handlers/            # HTTP handler(s)
│   │   └── logs.go
│   ├── models/              # Data models (LogEntry, etc.)
│   │   └── logentry.go
│   ├── storage/             # Persistence layer (file, DB)
│   │   └── storage.go
│   └── queue/               # Async queue implementation
│       └── queue.go
├── pkg/                     # Public libraries (if any)
│   └── util/                # Utility functions
│       └── timefmt.go
├── configs/                 # Configuration files (YAML/JSON/env)
│   └── config.yaml
├── scripts/                 # Helper scripts (migrations, setup)
│   └── migrate.sh
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
