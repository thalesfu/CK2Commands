package naimisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗企摩城LakhimpurBarony struct {
	feud.BaseBarony
}

var BLakhimpur罗企摩城 feud.Barony = &罗企摩城LakhimpurBarony{}

func init() {
	f := BLakhimpur罗企摩城.(*罗企摩城LakhimpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lakhimpur",
		TitleName: "罗企摩城",
		TitleCode: "b_lakhimpur",
	}
}
