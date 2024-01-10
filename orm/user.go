package orm

type User struct {
	Id   int
	Name string
	Age  string
}

func (User) TableName() string {
	return "user"
}
