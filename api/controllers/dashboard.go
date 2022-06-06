package api

import (
	"fmt"
	"net/http"
	"os"

	"Baco/POC-api/db"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDashboardAPI(c echo.Context) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials("AKID", os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY")),
	})
	fmt.Print(sess)
	fmt.Print(err)
	var conn *mongo.Client
	conn = db.ConnectToDB()
	fmt.Print(conn)
	// mongodb+srv://allensmilko:<password>
	// port := os.Getenv("PORT")
	// AWS_ACCESS_KEY
	// AWS_SECRET_KEY
	id := c.Param("t")
	return c.String(http.StatusOK, id)
}
