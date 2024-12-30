package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(data interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

type signUpReq struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConformPassword string `json:"conform_password"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	c := web.NewContext(w, r)
	req := &signUpReq{}
	err := c.ReadJson(req)
	if err != nil {
		_ = c.BadRequserJson(&commonResponse){
			BizCode: 4,
			Msg:fmt.Sprintf("invalid request: %v", err),
		}
		respBytes, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(respBytes))
		return
	}

	// do sign up
	fmt.Fprintf(w, "sign up success: %v", err)
}
