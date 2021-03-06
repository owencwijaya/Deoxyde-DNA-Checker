package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Time       string `json:"time"`
	Name       string `json:"name"`
	Disease    string `json:"disease"`
	Found      bool   `json:"found"`
	Similarity int    `json:"similarity"`
}

func match(c echo.Context) error {
	// if data yg diminta gada, return bad request
	r, _ := regexp.Compile(`^[CAGT]+$`)
	if c.FormValue("text") == "" || c.FormValue("disease") == "" || c.FormValue("method") == "" || !r.MatchString(c.FormValue("text")) {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
		c.Response().WriteHeader(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, "Text invalid or disease/method is empty")
	} else {
		//olah data
		text := c.FormValue("text")
		disease := c.FormValue("disease")
		db, err := sql.Open("mysql", databaseReference)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		_, err = db.Exec("USE stima3;")
		if err != nil {
			panic(err)
		}

		var pattern string

		if err := db.QueryRow("SELECT rantai FROM penyakit WHERE nama = ?", disease).Scan(&pattern); err != nil {
			panic(err)
		}

		var max_sim int
		if c.FormValue("method") == "BM" {
			fmt.Println("BM")
			max_sim = BMAlgo(text, pattern)
		} else {
			fmt.Println("KMP")
			max_sim = KMPAlgo(text, pattern)
		}
		fmt.Println(max_sim)

		percentage := int(float64(float64(max_sim)/float64(len(pattern))) * 100)

		res := &Response{
			Time:       time.Now().Local().Format("2006-01-02"),
			Name:       c.FormValue("username"),
			Disease:    c.FormValue("disease"),
			Found:      percentage >= 80,
			Similarity: percentage,
		}
		found := 0
		if res.Found {
			found = 1
		}
		_, err = db.Exec("INSERT INTO hasil_prediksi (tanggal, nama_pasien, penyakit_prediksi, status, kesamaan) VALUES (?, ?, ?, ?, ?);", res.Time, res.Name, res.Disease, found, res.Similarity)
		if err != nil {
			panic(err)
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
		c.Response().WriteHeader(http.StatusOK)
		return c.JSON(http.StatusOK, res)
	}
}
