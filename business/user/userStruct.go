package user

// omitempty 标签可让json序列化器忽略空值。
// user 类型
type UserStruct struct {
	Uuid     string `json:"uuid",db:"uuid"`
	Account  string `json:"account",db:"account"`
	Password string `json:"password",db:"password`
}
