package main 

type Tag struct {
	Id 	string
	Ts 	int
}

type TagVal struct {
	Tag 	Tag
	Key 	string
	Val		[]byte
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

func (t *Tag) update(id string) {
	t.Id = id
	t.Ts += 1
}