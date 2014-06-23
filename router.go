package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func init() {
	fmt.Println("router.go init()...")

	router.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "hello world")
	})

	router.HandleFunc("/goelia/subscribe/news/{weixinId}", func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		weixinId := vars["weixinId"]
		if FindWeixinId(weixinId) {
			m := make(map[string]interface{})
			m["weixinId"] = weixinId
			var u User
			if err := u.Get(weixinId); err != nil {
				fmt.Println("error:", err.Error())
				m["user"] = nil
			} else {
				m["user"] = u
			}

			var d Discribes
			items := d.FindByWeixinId(weixinId)
			m["items"] = items
			r.HTML(rw, 200, "index", m)
		} else {
			r.JSON(rw, 404, "you are not the weixin user.")
		}

		// var u User
		// if err := u.GetByWeixinId(weixinId); err != nil {
		// 	r.HTML(rw, 200, "index", nil)
		// } else {
		// 	r.HTML(rw, 200, "index", u)
		// }
	}).Methods("get")
	router.HandleFunc("/goelia/subscribe/news/{weixinId}", func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		weixinId := vars["weixinId"]
		fmt.Println("create...")
		bytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("request body: ", string(bytes))
		var u User

		if err = json.Unmarshal(bytes, &u); err != nil {
			panic(err)
		}
		fmt.Println("u:", u, u.WeixinId == "")
		if u.WeixinId == "" {
			u.Create(weixinId)
		} else {
			u.Update()
		}

	}).Methods("post")
	router.HandleFunc("/subscribe/items", func(rw http.ResponseWriter, req *http.Request) {
		var d DiscribeItem
		items := d.All()
		r.JSON(rw, 200, items)
		// if bytes, err := json.Marshal(items); err != nil {
		// 	panic(err)
		// } else {
		// 	r.JSON(rw, 200, bytes)
		// }
	}).Methods("get")
	router.HandleFunc("/subscribe/user/{weixin_id}", func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		weixinId := vars["weixin_id"]
		fmt.Println("weixinId", weixinId)

		var d Discribes
		items := d.FindByWeixinId(weixinId)
		r.JSON(rw, 200, items)
		// if bytes, err := json.Marshal(items); err != nil {
		// 	panic(err)
		// } else {
		// 	r.JSON(rw, 200, bytes)
		// }
	}).Methods("get")
	router.HandleFunc("/subscribe/users", func(rw http.ResponseWriter, req *http.Request) {
		var u User
		users := u.All()
		r.JSON(rw, 200, users)
	}).Methods("get")
	router.HandleFunc("/subscribe/create", func(rw http.ResponseWriter, req *http.Request) {
		bytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("request body: ", string(bytes))
		var d DiscribeItem

		if err = json.Unmarshal(bytes, &d); err != nil {
			panic(err)
		}
		fmt.Println("d:", d)
		d.Create()
	}).Methods("post")
	router.HandleFunc("/subscribe/{id}", func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		fmt.Println("id", id)
		var d DiscribeItem
		d.Delete(id)
	}).Methods("delete")

	router.HandleFunc("/admin/user", func(rw http.ResponseWriter, req *http.Request) {
		// fmt.Println("create...")
		// bytes, err := ioutil.ReadAll(req.Body)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println("request body: ", string(bytes))
		// var u User

		// if err = json.Unmarshal(bytes, &u); err != nil {
		// 	panic(err)
		// }
		// fmt.Println("u:", u)
		// if err = u.Create(); err != nil {
		// 	panic(err)
		// } else {
		// 	r.JSON(rw, 200, "success")
		// }
		r.HTML(rw, 200, "admin-user", nil)
	}).Methods("get")

	router.HandleFunc("/create", func(rw http.ResponseWriter, req *http.Request) {
		// fmt.Println("create...")
		// bytes, err := ioutil.ReadAll(req.Body)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println("request body: ", string(bytes))
		// var u User

		// if err = json.Unmarshal(bytes, &u); err != nil {
		// 	panic(err)
		// }
		// fmt.Println("u:", u)
		// if err = u.Create(); err != nil {
		// 	panic(err)
		// } else {
		// 	r.JSON(rw, 200, "success")
		// }
	}).Methods("post")

	router.HandleFunc("/admin", func(rw http.ResponseWriter, req *http.Request) {
		r.HTML(rw, 200, "admin", nil)
	}).Methods("get")
}
