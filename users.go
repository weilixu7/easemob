// Copyright 2015 The Yang Gui AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package easemob

import "fmt"

// UsersService handles communication with the user related
// methods of the Easemob API.
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/
type UsersService struct {
	client *Client
}

// Register a user without Authorization
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#im
func (s *UsersService) RegisterWithoutAuth(username string, password string, nickname string) (*Response, error) {
	put := &PutOptions{Username: username, Password: password, Nickname: nickname}

	var u string
	u = "users"

	req, err := s.client.NewRequestWithoutAuth("POST", u, put)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// Register users with Authorization
func (s *UsersService) Registers(usernames []string, password string) (*Response, error) {
	puts := []PutOptions{}
	for _, username := range usernames {
		put := PutOptions{Username: username, Password: password}
		puts = append(puts, put)
	}

	var u string
	u = "users"

	req, err := s.client.NewRequest("POST", u, puts)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// Register a user with Authorization
func (s *UsersService) Register(username string, password string) (*Response, error) {
	put := &PutOptions{Username: username, Password: password}

	var u string
	u = "users"

	req, err := s.client.NewRequest("POST", u, put)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// group Register a user with Authorization
func (s *UsersService) RegisterGroup(users map[string]string) (*Response, error) {
	puts := []PutOptions{}
	for username, password := range users {
		put := PutOptions{Username: username, Password: password}
		puts = append(puts, put)
	}

	var u string
	u = "users"

	req, err := s.client.NewRequest("POST", u, puts)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// user status online offline
func (s *UsersService) UserStatus(username string) (*Response, error) {
	put := &PutOptions{}

	var u string
	u = "users/" + username + "/status"

	req, err := s.client.NewRequest("GET", u, put)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// Disconnect user
func (s *UsersService) Disconnect(username string) (*Response, error) {
	put := &PutOptions{}

	var u string
	u = "users/" + username + "/disconnect"

	req, err := s.client.NewRequest("GET", u, put)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// Get fetches a User.
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#im-2
func (s *UsersService) Get(owner string) (*Response, error) {
	var u string
	u = fmt.Sprintf("users/%v", owner)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// ListAll lists all Easemob users.
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#im-3
func (s *UsersService) ListAll(opt *ListOptions) (*Response, error) {
	u, err := addOptions("users", opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

// Delete a user.
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#im-5
func (s *UsersService) Delete(owner string) (*Response, error) {
	var u string
	u = fmt.Sprintf("users/%v", owner)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req)
}

// ResetPassword
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#resetpassword
func (s *UsersService) ResetPassword(owner string, password string) (*Response, error) {
	var u string
	u = fmt.Sprintf("users/%v/password", owner)

	opt := &PutOptions{NewPass: password}
	req, err := s.client.NewRequest("PUT", u, opt)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// EditUsername
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#nickname
func (s *UsersService) EditNickname(owner string, nickname string) (*Response, error) {
	var u string
	u = fmt.Sprintf("users/%v", owner)

	opt := &PutOptions{Nickname: nickname}
	req, err := s.client.NewRequest("PUT", u, opt)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// AddFriend add a friend.
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#contacts
func (s *UsersService) AddFriend(owner string, friend string) (*Response, error) {
	var u string
	u = fmt.Sprintf("users/%v/contacts/users/%v", owner, friend)

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// DeleteFriend
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#delfriend
func (s *UsersService) DeleteFriend(owner string, friend string) (*Response, error) {
	var u string
	u = fmt.Sprintf("users/%v/contacts/users/%v", owner, friend)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// DeleteFriends
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#im-6
// func (s *UsersService) DeleteFriends(start time.Time, end time.Time) (*Response, error) {
//     var u string
//     u = "users"

//     // opt := &ListOptions{QL: fmt.Sprintf("created>%d", ...)}

// }

// GetFriends
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#queryfriend
func (s *UsersService) GetFriends(owner string) (*Response, error) {
	var u string
	u = fmt.Sprintf("users/%v/contacts/users", owner)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// GetBlocks
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#blocksusers
func (s *UsersService) GetBlocks(owner string) (*Response, error) {
	var u string
	u = fmt.Sprintf("users/%v/blocks/users", owner)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// AddBlocks
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#addblocksusers
func (s *UsersService) AddBlocks(owner string, usernames []string) (*Response, error) {
	var u string
	u = fmt.Sprintf("users/%v/blocks/users", owner)

	opt := &PutOptions{Usernames: usernames}
	req, err := s.client.NewRequest("POST", u, opt)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// DeleteBlock
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#delblocksusers
func (s *UsersService) DeleteBlock(owner string, friend string) (*Response, error) {
	var u string
	u = fmt.Sprintf("users/%v/blocks/users/%v", owner, friend)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// Status
//
// Easemob API docs: http://www.easemob.com/docs/rest/sendmessage/#status
func (s *UsersService) Status(owner string) (*Response, error) {
	var u string
	u = fmt.Sprintf("users/%v/status", owner)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// UserGroups
//
// Easemob API docs: http://www.easemob.com/docs/rest/groups/#joinedchatgroups
func (s *GroupService) UserGroups(user string) (*Response, error) {
	var u string
	u = fmt.Sprintf("users/%v/joined_chatgroups", user)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}

// Offline Message Count
//
// Easemob API docs: http://www.easemob.com/docs/rest/userapi/#msgcount
func (s *UsersService) OfflineMsgCount(username string) (*Response, error) {

	var (
		url string
	)

	url = fmt.Sprintf("users/%v/offline_msg_count", username)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	return resp, err
}
