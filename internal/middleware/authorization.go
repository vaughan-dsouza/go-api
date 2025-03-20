package middleware

import(
	"errors"
	"net/http"
	"github.com/vaughan-dsouza/go-api/api"
	"github.com/vaughan-dsouza/go-api/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("invalid username or token.")

func Authorization(next http.Handler)http.Handler {
	return http.HandlerFunc(func( w http.ResponseWriter, r *http.Request){
		var username string = r.URL.Query().Get("username")
		var token =  r.Header.Get("Authorization")
		var err error 

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil{
			api.internalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetials(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)){
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		} 

		next.ServeHTTP(w,r)
	})
}