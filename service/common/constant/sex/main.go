package sex

/*
性别
*/
type Sex uint64

func (s Sex) String() string {
	switch s {
	case SexMan:
		return "男"
	case SexWoman:
		return "女"
	default:
		return "未知"
	}
}

const (
	SexUnknown Sex = iota
	SexMan
	SexWoman
)
