package main

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type string
	MessageCharLimit int
}

const (
    MembershipTypePremium = "premium"
    MembershipTypeBasic   = "standard"
)

func newUser(name string, membershipType string) User {
    var membership Membership

    if membershipType == MembershipTypePremium {
        membership.Type = MembershipTypePremium
        membership.MessageCharLimit = 1000
    } else {
        membership.Type = MembershipTypeBasic
        membership.MessageCharLimit = 100
    }

    user := User{
        Name:       name,
        Membership: membership,
    }

    return user
}
