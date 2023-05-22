module {{ .ProjectName }}

go 1.19

require (
	github.com/golang/mock v1.6.0
	github.com/gustavohenrique/coolconf v1.0.2
	github.com/sirupsen/logrus v1.8.1
  {{ if .HasDgraph }}
	github.com/dgraph-io/dgo/v210 v210.0.0-20220113041351-ba0e5dfc4c3e
	github.com/json-iterator/go v1.1.12
  {{ end }}
  {{ if .HasPostgres }}
	github.com/jackc/pgx/v4 v4.16.1
	github.com/lib/pq v1.10.6
  {{ end }}
  {{ if .HasSqlite }}
	github.com/mattn/go-sqlite3 v1.14.13
  {{ end }}
  {{ if or .HasPostgres .HasSqlite }}
	github.com/jmoiron/sqlx v1.3.5
  {{ end }}
  {{ if .HasHttpServer }}
	github.com/labstack/echo/v4 v4.7.2
	github.com/swaggo/echo-swagger v1.3.0
	github.com/swaggo/swag v1.8.4
  {{ end }}
  {{ if .HasGrpcWebServer }}
	github.com/improbable-eng/grpc-web v0.15.0
  {{ end }}
  {{ if .HasGrpcServer }}
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
  {{ end }}
  {{ if .HasNatsServer }}
	github.com/nats-io/nats-server/v2 v2.1.2
	github.com/nats-io/nats.go v1.9.1
  {{ end }}
)
