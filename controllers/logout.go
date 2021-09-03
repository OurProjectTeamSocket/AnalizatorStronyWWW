package adds

import "net/http"

func Logout() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		c := &http.Cookie{
			Name:     "id",
			MaxAge:   -1, // Ciasteczko nie istniejące
			HttpOnly: true,
		}

		http.SetCookie(w, c)

		req.AddCookie(c)

		http.Redirect(w, req,"/", http.StatusSeeOther)
	})
}