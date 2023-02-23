package laodikeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 老底嘉LaodikeiaBarony struct {
	feud.BaseBarony
}

var BLaodikeia老底嘉 feud.Barony = &老底嘉LaodikeiaBarony{}

func init() {
	f := BLaodikeia老底嘉.(*老底嘉LaodikeiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laodikeia",
		TitleName: "老底嘉",
		TitleCode: "b_laodikeia",
	}
}
