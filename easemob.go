// Copyright 2015 The go-easemob AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package easemob

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"gitlab.xiaoenai.net/server/log"

	"github.com/google/go-querystring/query"
)

const (
	debug        = false
	repeat_times = 3
	baseURL      = "https://a1.easemob.com/"
	grantTYPE    = "client_credentials"
	mediaType    = "application/json"
)

var (
	repeat = 0
)

// A Client manages communication with the Easemob API
type Client struct {
	// HTTP Client used to communicate with the API.
	client *http.Client

	// Base URL for API requeest.
	BaseURL *url.URL

	// Easemob Org and App info.
	OrgName string
	AppName string

	// Credential
	Credentials *Credentials

	// Authorization
	Token   string
	Expires int64

	Users    *UsersService
	Messages *MessagesService
	Groups   *GroupService
}

// ListOptions specifies the optional parameters to various list methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Limit  int    `url:"limit,omitempty"`
	Cursor string `url:"cursor,omitempty"`
	QL     string `url:"ql,omitempty"`
}

// PutOptions specifies the parameters to various put methods.
type PutOptions struct {
	Username    string   `json:"username,omitempty"`
	Nickname    string   `json:"nickname,omitempty"`
	Password    string   `json:"password,omitempty"`
	NewPass     string   `json:"newpassword,omitempty"`
	Usernames   []string `json:"usernames,omitempty"`
	Groupname   string   `json:"groupname,omitempty"`
	Description string   `json:"description,omitempty"`
	Maxusers    int      `json:"maxusers,omitempty"`
}

type MessagePutOptions struct {

	/**
	 * :param TargetType: users 给用户发消息, chatgroups 给群发消息
	 *
	 * :param Target: 注意这里需要用数组, 数组长度建议不大于20, 即使只有一个用户,
	 *                也要用数组 ['u1'], 给用户发送时数组元素是用户名,
	 *                给群组发送时数组元素是groupid
	 *
	 * :param Message:  消息内容，参考[聊天记录]
	 *                  (http://www.easemob.com/docs/rest/chatmessage/)
	 *                  里的bodies 内容
	 *
	 * :param From: 表示这个消息是谁发出来的, 可以没有这个属性, 那么就会显示是admin,
	 *              如果有的话, 则会显示是这个用户发出的。
	 *
	 * :param Extend: 扩展属性, 由app自己定义.可以没有这个字段，但是如果有，
	 *                值不能是“ext:null“这种形式，否则出错
	 */

	TargetType string            `json:"target_type"`
	Target     []string          `json:"target"`
	Msg        *MessageType      `json:"msg"`
	From       string            `json:"from,omitempty"`
	Extend     map[string]string `json:"ext,omitempty"`
}

type MessageType struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

// addOptions adds the parameters in opt as URL query parameters to s.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

type Credentials struct {
	GrantType string `json:"grant_type"`
	ClientId  string `json:"client_id"`
	Secret    string `json:"client_secret"`
}

// NewClient returns a new Easemob API client
func NewClient(clientId string, clientSecret string, orgName string, appName string, token string) (*Client, error) {
	c := new(Client)
	c.client = http.DefaultClient
	c.BaseURL, _ = url.Parse(baseURL)
	c.OrgName = orgName
	c.AppName = appName
	c.Credentials = &Credentials{
		GrantType: grantTYPE,
		ClientId:  clientId,
		Secret:    clientSecret}

	c.Token = token

	c.Users = &UsersService{client: c}
	c.Messages = &MessagesService{client: c}
	c.Groups = &GroupService{client: c}

	return c, nil
}

// Response is a Easemob API response.
type Response struct {
	*http.Response

	AccessToken string `json:"access_token,omitempty"`
	Expires     int64  `json:"expires_in,omitempty"`
	Application string `json:"application,omitempty"`

	Action string `json:"action,omitempty"`
	Params struct {
		Limit  []string
		Cursor []string
	} `json:"params,omitempty"`
	Path             string  `json:"path,omitempty"`
	URI              string  `json:"uri,omitempty"`
	Timestamp        int64   `json:"timestamp,omitempty"`
	Duration         int     `json:"duration,omitempty"`
	Organization     string  `json:"organization,omitempty"`
	ApplicationName  string  `json:"applicationName,omitempty"`
	Cursor           string  `json:"cursor,omitempty"`
	Count            int     `json:"count,omitempty"`
	Entities         []*User `json:"entities,omitempty"`
	Error            string  `json:"error,omitempty"`
	Exception        string  `json:"exception,omitempty"`
	ErrorDescription string  `json:"error_description,omitempty"`
	// Data             []string `json:"data,omitempty"`
}

type User struct {
	Uuid                     string `json:"uuid"`
	Type                     string `json:"user"`
	Created                  int64  `json:"created"`
	Modified                 int64  `json:"modified"`
	Username                 string `json:"username"`
	Activated                bool   `json:"activated"`
	Nickname                 string `json:"nickname"`
	NotifierName             string `json:"notifier_name"`
	NotificationDisplayStyle int    `json:"notification_display_style"`
	NotificationNoDisturing  bool   `json:"notification_no_disturbing"`
}

// NewRequest creates an API request. A relative URL can be provided in urlStr
// in which case it is resolved relative to the BaesURL of the Client.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	req, err := c.buildRequest(method, urlStr, body)
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", "Bearer", c.Token))
	return req, err
}

func (c *Client) NewRequestWithoutAuth(method, urlStr string, body interface{}) (*http.Request, error) {
	return c.buildRequest(method, urlStr, body)
}

func (c *Client) buildRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(fmt.Sprintf("%v/%v/%v", c.OrgName, c.AppName, urlStr))
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request) (*Response, error) {

	resp, err := c.client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	code := resp.StatusCode
	if code == 408 {
		//timeout repeat req
		if repeat < repeat_times {
			repeat++
			return c.Do(req)
		}
	} else if code == 503 {
		//limit req
		time.Sleep(500 * time.Millisecond)
		if repeat < repeat_times {
			repeat++
			return c.Do(req)
		}
	}
	err = CheckResponse(resp)

	response := new(Response)
	json.Unmarshal(body, response)
	repeat = 0
	response.Response = resp
	return response, err
}

// Token get auth token
func (c *Client) GetToken() error {
	var u string
	u = "token"

	req, err := c.NewRequestWithoutAuth("POST", u, c.Credentials)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	c.Token = resp.AccessToken
	c.Expires = resp.Expires

	log.Tracef("token:%v,expires:%v", c.Token, c.Expires)

	return nil
}

// An ErrorResponse reports one or more errors caused by an API request.
// Easemob docs: http://www.easemob.com/docs/helps/errorcodes/
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

// CheckResponse checks the API response for errors, and returns them if
// present.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}
