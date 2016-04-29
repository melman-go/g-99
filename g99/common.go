package g99

type G99Error struct {
	error
	Msg string
}

func (err G99Error)Error() string {
	return err.Msg
}

func GetListLen(list []interface{}) (int, error) {
	if list==nil {
		return 0, G99Error{Msg:"list is nil"}
	}
	len := len(list)
	if len==0 { return 0, G99Error{Msg:"list is empty"}}
	return len, nil
}