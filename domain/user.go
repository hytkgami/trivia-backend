package domain

type User struct {
	UID  string `db:"uid"`
	Name string `db:"name"`
}
