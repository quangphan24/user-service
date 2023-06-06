package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strings"
	"user-service/conf"
	authenpb "user-service/proto/authen"
)

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Request().Header.Get("Authorization")
		data := strings.Split(token, " ")[1]
		conn, err := grpc.Dial(conf.GetConfig().GRPCServer.AuthenServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logrus.Error("did not connect: %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		defer conn.Close()
		client := authenpb.NewAuthenServiceClient(conn)
		_, err = client.VerifyToken(context.Background(), &authenpb.String{Value: data})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return next(c)
	}
}
