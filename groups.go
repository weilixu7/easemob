package easemob

import (
  "fmt"
)

// GroupService handles communication with the group related
// methods of the Easemob API.
//
// Easemob API docs: http://www.easemob.com/docs/rest/groups/
type GroupService struct {
  client *Client
}

// ListAll list all groups
//
// Easemob API docs: http://www.easemob.com/docs/rest/groups/#getallgroups
func (s *GroupService) ListAll() (*Response, error) {
  var u string
  u = "chatgroups"
  req, err := s.client.NewRequest("GET", u, nil)
  if err != nil {
    return nil, err
  }

  resp, err := s.client.Do(req)
  return resp, err
}

func generalizeStringList(strs []string) string {
  var strings string
  strings = ""
  for i := range strs {
    strings = fmt.Sprintf("%v,%v", strings, strs[i])
  }
  return strings
}

// Get fetch group details
//
// Easemob API docs: http://www.easemob.com/docs/rest/groups/#getgroups
func (s *GroupService) Get(groups ...string) (*Response, error) {
  var u string
  u = fmt.Sprintf("chatgroups/%v", generalizeStringList(groups))

  req, err := s.client.NewRequest("GET", u, nil)
  if err != nil {
    return nil, err
  }

  resp, err := s.client.Do(req)
  return resp, err
}

// Create create a new group
//
// Easemob API docs: http://www.easemob.com/docs/rest/groups/#create
func (s *GroupService) Create() (*Response, error) {
  var u string
  u = "chatgroups"

  req, err := s.client.NewRequest("POST", u, nil)
  if err != nil {
    return nil, err
  }

  resp, err := s.client.Do(req)
  return resp, err
}

// Update edit a group infomation
//
// Easemob API docs: http://www.easemob.com/docs/rest/groups/#update
func (s *GroupService) Update(groupid string, name string, description string, maxusers int) (*Response, error) {
  var u string
  u = fmt.Sprintf("chatgroups/%v", groupid)

  put := new(PutOptions)
  if name != "" {
    put.Groupname = name
  }
  if description != "" {
    put.Description = description
  }
  if maxusers > 0 {
    put.Maxusers = maxusers
  }

  req, err := s.client.NewRequest("PUT", u, put)
  if err != nil {
    return nil, err
  }

  resp, err := s.client.Do(req)
  return resp, err
}

// Delete remove a group
//
// Easemob API docs: http://www.easemob.com/docs/rest/groups/#delete
func (s *GroupService) Delete(groupid string) (*Response, error) {
  var u string
  u = fmt.Sprintf("chatgroups/%v", groupid)

  req, err := s.client.NewRequest("DELETE", u, nil)
  if err != nil {
    return nil, err
  }

  resp, err := s.client.Do(req)
  return resp, err
}

// Members
//
// Easemob API docs: http://www.easemob.com/docs/rest/groups/#users
func (s *GroupService) Members(groupid string) (*Response, error) {
  var u string
  u = fmt.Sprintf("chatgroups/%v/users", groupid)

  req, err := s.client.NewRequest("GET", u, nil)
  if err != nil {
    return nil, err
  }

  resp, err := s.client.Do(req)
  return resp, err
}

// AddMember
//
// Easemob API docs: http://www.easemob.com/docs/rest/groups/#addmember
func (s *GroupService) AddMember(groupid string, user string) (*Response, error) {
  var u string
  u = fmt.Sprintf("chatgroups/%v/users/%v", groupid, user)

  req, err := s.client.NewRequest("POST", u, nil)
  if err != nil {
    return nil, err
  }

  resp, err := s.client.Do(req)
  return resp, err
}

// DeleteMember
//
// Easemob API docs: http://www.easemob.com/docs/rest/groups/#deletemember
func (s *GroupService) DeleteMember(groupid string, user string) (*Response, error) {
  var u string
  u = fmt.Sprintf("chatgroups/%v/users/%v", groupid, user)

  req, err := s.client.NewRequest("DELETE", u, nil)
  if err != nil {
    return nil, err
  }

  resp, err := s.client.Do(req)
  return resp, err
}

// AddMembers
//
// Easemob API docs: http://www.easemob.com/docs/rest/groups/#addmemberbatch
func (s *GroupService) AddMembers(users ...string) (*Response, error) {
  var u string
  u = fmt.Sprintf("chatgroups/%v/users")

  put := &PutOptions{Usernames: users}
  req, err := s.client.NewRequest("POST", u, put)
  if err != nil {
    return nil, err
  }

  resp, err := s.client.Do(req)
  return resp, err
}
