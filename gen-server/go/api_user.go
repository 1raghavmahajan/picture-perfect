/*
 * Picture Perfect
 *
 * Sample API description. You can find more at.
 *
 * API version: 0.1.0
 * Contact: 1raghavmahajan@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"net/http"
)

func AddReviewToMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
