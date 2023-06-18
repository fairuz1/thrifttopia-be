package middleware

import (
	"net/http"
	"strings"
	"thriftopia/helper"

	"github.com/dgrijalva/jwt-go"
)

var ResponseError = helper.ResponseError

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Baca token dari header Authorization
		tokenString := r.Header.Get("Authorization")

		// Periksa apakah token ada dan valid
		if tokenString == "" {
			ResponseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Split header Authorization dan ambil token bagian kedua setelah "Bearer"
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 {
			ResponseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		token := parts[1]

		// Verifikasi token menggunakan secret key yang sama dengan yang digunakan saat pembuatan token
		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret-key-thriftopia"), nil
		})

		if err != nil {
			ResponseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Token valid, lanjutkan dengan permintaan berikutnya
		next.ServeHTTP(w, r)
	})
}

func AdminAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Baca token dari header Authorization
		tokenString := r.Header.Get("Authorization")

		// Periksa apakah token ada dan valid
		if tokenString == "" {
			ResponseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Split header Authorization dan ambil token bagian kedua setelah "Bearer"
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ResponseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		tokenJWT := parts[1]

		// Verifikasi token menggunakan secret key yang sama dengan yang digunakan saat pembuatan token
		token, err := jwt.Parse(tokenJWT, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret-key-thriftopia"), nil
		})

		if err != nil {
			ResponseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// verifikasi token
		if token.Valid {
			// Memeriksa peran pengguna dari token
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			role, ok := claims["role"].(string)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Memeriksa otorisasi berdasarkan peran
			if role == "admin" {
				// Jika pengguna adalah admin, lanjutkan ke handler selanjutnya
				next.ServeHTTP(w, r)
			} else {
				// Jika pengguna bukan admin, berikan respons akses ditolak
				w.WriteHeader(http.StatusForbidden)
				return
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// next.ServeHTTP(w, r)
	})
}
