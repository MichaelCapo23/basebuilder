package user

type Service interface {
	UserProfile
}

type UserProfile interface {
	HandleGetProfile()
}
