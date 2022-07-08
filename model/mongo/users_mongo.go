package mongo

type UsersMongo struct {
}

func (UsersMongo) TableName() string {
	return "users"
}
