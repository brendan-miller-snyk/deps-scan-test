package main

import (
	"fmt"

	_ "cloud.google.com/go/bigquery"
	_ "cloud.google.com/go/cloudtasks/apiv2beta3"
	_ "cloud.google.com/go/datastore"
	_ "cloud.google.com/go/iam/credentials/apiv1"
	_ "cloud.google.com/go/logging"
	_ "cloud.google.com/go/storage"
	_ "contrib.go.opencensus.io/exporter/stackdriver"
	_ "contrib.go.opencensus.io/exporter/stackdriver/propagation"
	_ "github.com/gin-contrib/cors"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-playground/mold/v4/modifiers"
	_ "github.com/go-playground/validator/v10"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/google/uuid"
	_ "github.com/lib/pq"
	_ "github.com/r3labs/diff"
	_ "github.com/sendgrid/sendgrid-go"
	_ "github.com/sendgrid/sendgrid-go/helpers/mail"
	_ "github.com/tent/hawk-go"
	_ "go.opencensus.io/plugin/ochttp"
	_ "go.opencensus.io/trace"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	_ "golang.org/x/oauth2"
	_ "golang.org/x/oauth2/google"
	_ "google.golang.org/api/googleapi"
	_ "google.golang.org/api/iterator"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	_ "google.golang.org/genproto/googleapis/cloud/tasks/v2beta3"
	_ "google.golang.org/genproto/googleapis/iam/credentials/v1"
	_ "google.golang.org/genproto/googleapis/api/monitoredres"
)

func main() {
	fm.Println("ok")
}
