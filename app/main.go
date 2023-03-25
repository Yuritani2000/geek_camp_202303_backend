package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Data struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type Trip struct {
	Id             int    `json:"id"`
	HostId         string `json:"host_id"`
	HostName       string `json:"host_name"`
	CarLicense     string `json:"car_license"`
	CarName        string `json:"car_name"`
	PassengerLimit int    `json:"passenger_limit"`
	LocationFrom   string `json:"location_from"`
	LocationTo     string `json:"location_to"`
}

type Request struct {
	Id            int    `json:"id"`
	TripId        int    `json:"trip_id"`
	PassengerId   string `json:"passenger_id"`
	Passengername string `json:"passenger_name"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {

	// datas := make([]*Data, 0, 100)
	trips := make([]*Trip, 0, 100)
	requests := make([]*Request, 0, 100)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// トリップを新規追加
	r.Post("/trip", func(w http.ResponseWriter, r *http.Request) {
		var req Trip

		log.Println(r.Body)

		// JSONデータをデコードしてデータの構造体に変換している
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// 変更できなかったときにはエラーを返すようにしている．
			body := &Error{
				Code:    400,
				Message: "Bad request",
			}
			resBody, err := json.Marshal(body)
			if err != nil {
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
			w.WriteHeader(400)
			w.Write(resBody)
			return
		}

		// 成功したときにはdatasに値を格納
		trips = append(trips, &req)
		log.Println(trips)

		resBody, _ := json.Marshal(req)

		// 成功した場合は，BODYに受け取ったリクエストを返すのがサーバの流儀
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
		w.WriteHeader(http.StatusOK)
		w.Write(resBody)
	})

	// トリップ一覧を取得
	r.Get("/trip", func(w http.ResponseWriter, r *http.Request) {
		// スライスに格納されているデータ(キャッシュされているデータ)をresに代入
		res := trips

		// POSTのときと逆で，デコードしていたデータをJSONとしてエンコードしている．
		resBody, err := json.Marshal(res)
		if err != nil {
			body := &Error{
				Code:    500,
				Message: "internal server error",
			}
			resBody, err := json.Marshal(body)
			if err != nil {
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
			w.WriteHeader(500)
			w.Write(resBody)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
		w.WriteHeader(http.StatusOK)
		w.Write(resBody)
	})

	// 乗車リクエストを新規追加
	r.Post("/request", func(w http.ResponseWriter, r *http.Request) {
		var req Request

		// JSONデータをデコードしてデータの構造体に変換している
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// 変更できなかったときにはエラーを返すようにしている．
			body := &Error{
				Code:    400,
				Message: "Bad request",
			}
			resBody, err := json.Marshal(body)
			if err != nil {
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
			w.WriteHeader(400)
			w.Write(resBody)
			return
		}

		// 成功したときにはdatasに値を格納
		requests = append(requests, &req)
		log.Println(requests)

		resBody, _ := json.Marshal(req)

		// 成功した場合は，BODYに受け取ったリクエストを返すのがサーバの流儀
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
		w.WriteHeader(http.StatusOK)
		w.Write(resBody)
	})

	http.ListenAndServe(":8080", r)
}
