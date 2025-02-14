package casbin

import (
	"net/http"

	"github.com/cgisoftware/initializers/auth"
)

func (casbinConfig *CasbinConfig) Middleware(contextValue auth.ContextValue) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			user := auth.GetStringFromContext(request.Context(), contextValue)
			if user == "" {
				http.Error(response, "Unauthorized", http.StatusUnauthorized)
				return
			}

			path := request.URL.Path
			method := request.Method

			allowed, err := casbinConfig.Enforce.Enforce(user, path, method)
			if err != nil {
				http.Error(response, "Error enforcing policy", http.StatusInternalServerError)
				return
			}

			if !allowed {
				http.Error(response, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(response, request)
		})
	}
}
