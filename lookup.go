// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package user

// Current returns the current user.
func Current() (*User, error) {
	return current()
}

// Lookup looks up a user by username. If the user cannot be found, the
// returned error is of type UnknownUserError.
func Lookup(username string) (*User, error) {
	return lookupUser(username)
}

// LookupId looks up a user by userid. If the user cannot be found, the
// returned error is of type UnknownUserIdError.
func LookupId(uid string) (*User, error) {
	return lookupUserId(uid)
}

// CurrentGroup returns the current group.
func CurrentGroup() (*Group, error) {
	return currentGroup()
}

// LookupGroup looks up a group by name. If the group cannot be found, the
// returned error is of type UnknownGroupError.
func LookupGroup(groupname string) (*Group, error) {
	return lookupGroup(groupname)
}

// LookupGroupId looks up a group by groupid. If the group cannot be found, the
// returned error is of type UnknownGroupIdError.
func LookupGroupId(gid string) (*Group, error) {
	return lookupGroupId(gid)
}

// In indicates whether the user is a member of the given group.
func (u *User) In(g *Group) (bool, error) {
	return userInGroup(u, g)
}

// Members returns the list of members of the group.
func (g *Group) Members() ([]string, error) {
	return groupMembers(g)
}
