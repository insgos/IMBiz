package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	imdb "imbiz/internal/imdb"
	users "imbiz/internal/users"

	"github.com/labstack/echo/v4"
)

// @Summary Read User
// @Description Read User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /userinfo/read [get]
func ReadUser(c echo.Context) error {
	var buff []byte
	var e error
	var u users.User

	db := imdb.NewDb()
	defer imdb.CloseDb(db)

	id := c.FormValue("userid")
	pw := c.FormValue("pw")

	rows, e := db.Query("SELECT USER_ID, PW, NAME, EMAIL, NUMBER, DEPARTMENT_ID, POSITION, LAST_ACCESS_DT, UPDATE_DT, CREATE_DT FROM USER where USER_ID = ?", id)
	if e != nil {
		return e
	}
	defer rows.Close()

	for rows.Next() {
		e := rows.Scan(&u.UserId, &u.Pw, &u.Name, &u.Email, &u.Number, &u.DepartmentId, &u.Position, &u.LastAccessDt, &u.UpdateDt, &u.CreateDt)
		if e != nil {
			return e
		}
	}

	if id == u.UserId && pw == u.Pw {
		buff, e = json.Marshal(u)
		if e != nil {
			return e
		}
	}

	return c.String(http.StatusOK, string(buff))
}

// @Summary  Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /userinfo/create [post]
func CreateUser(c echo.Context) error {
	db := imdb.NewDb()
	defer imdb.CloseDb(db)

	u := &users.User{
		UserId:       c.FormValue("userid"),
		Pw:           c.FormValue("pw"),
		Name:         c.FormValue("name"),
		Email:        c.FormValue("email"),
		Number:       c.FormValue("number"),
		DepartmentId: c.FormValue("departmentid"),
		Position:     c.FormValue("position"),
		LastAccessDt: time.Now(),
		UpdateDt:     time.Now(),
		CreateDt:     time.Now(),
	}

	_, e := db.Exec(`INSERT INTO USER(USER_ID, PW, NAME, EMAIL, NUMBER, DEPARTMENT_ID, POSITION, LAST_ACCESS_DT, UPDATE_DT, CREATE_DT) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		u.UserId,
		u.Pw,
		u.Name,
		u.Email,
		u.Number,
		u.DepartmentId,
		u.Position,
		u.LastAccessDt,
		u.UpdateDt,
		u.CreateDt,
	)
	if e != nil {
		log.Println(e)
	}
	return c.String(http.StatusOK, "success")
}

// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /userinfo/update [put]
func UpdateUser(c echo.Context) error {
	return nil
}

// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /userinfo/delete [delete]
func DeleteUser(c echo.Context) error {
	var e error
	var u users.User

	db := imdb.NewDb()
	defer imdb.CloseDb(db)

	id := c.FormValue("userid")
	pw := c.FormValue("pw")

	rows, e := db.Query("SELECT USER_ID, PW FROM USER WHERE USER_ID = ?", id)
	if e != nil {
		return e
	}
	defer rows.Close()

	for rows.Next() {
		e := rows.Scan(&u.UserId, &u.Pw)
		if e != nil {
			return e
		}
	}

	if id == u.UserId && pw == u.Pw {
		_, e = db.Exec("DELETE FROM USER WHERE USER_ID = ?", id)
		if e != nil {
			return e
		}
	}

	return c.String(http.StatusOK, "success")
}

// @Summary Login
// @Description Login
// @Tags Login
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /userinfo/login [POST]
func Login(c echo.Context) error {
	return c.String(http.StatusOK, "success")
}

// @Summary Logout
// @Description Logout
// @Tags Login
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /userinfo/logout [POST]
func Logout(c echo.Context) error {
	return c.String(http.StatusOK, "success")
}
