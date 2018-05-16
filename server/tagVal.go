package main 

type Tag struct {
	Id 	string
	Ts 	int
}

type TagVal struct {
	Tag 	Tag
	Key 	string
	Val		string
}

func (t *Tag) smaller(x Tag) bool {
	var res bool
	if t.Ts < x.Ts {
		res = true
	} else if t.Ts > x.Ts {
		res = false
	} else {
		res = t.Id < x.Id
	}
	return res
}

func (tv *TagVal) update(id string, val string) {
	tv.Tag.Id = id
	tv.Tag.Ts += 1
	tv.Val = val
}
