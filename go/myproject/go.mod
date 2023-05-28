module {{ .ProjectName }}

go 1.20

require (
	github.com/golang/mock v1.6.0
	github.com/gustavohenrique/coolconf v1.0.2
	github.com/sirupsen/logrus v1.9.2
  {{ if .HasDgraph }}
	github.com/dgraph-io/dgo/v210 v210.0.0-20230328113526-b66f8ae53a2d
	github.com/json-iterator/go v1.1.12
  {{ end }}
  {{ if .HasPostgres }}
	github.com/jackc/pgx/v4 v4.18.1
	github.com/lib/pq v1.10.9
  {{ end }}
  {{ if .HasSqlite }}
	github.com/mattn/go-sqlite3 v1.14.16
  {{ end }}
  {{ if or .HasPostgres .HasSqlite }}
	github.com/jmoiron/sqlx v1.3.5
  {{ end }}
  {{ if .HasHttpServer }}
	github.com/labstack/echo/v4 v4.10.2
	github.com/swaggo/echo-swagger v1.4.0
	github.com/swaggo/swag v1.16.1
  {{ end }}
  {{ if .HasGrpcWebServer }}
	github.com/improbable-eng/grpc-web v0.15.0
  {{ end }}
  {{ if .HasGrpcServer }}
	google.golang.org/protobuf v1.5.3
	google.golang.org/grpc v1.55.0
  {{ end }}
  {{ if .HasNatsServer }}
	github.com/nats-io/nats-server/v2 v2.9.17
	github.com/nats-io/nats.go v1.26.0
  {{ end }}
)
