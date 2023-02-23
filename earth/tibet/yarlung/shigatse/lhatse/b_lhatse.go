package lhatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉孜LhatseBarony struct {
	feud.BaseBarony
}

var BLhatse拉孜 feud.Barony = &拉孜LhatseBarony{}

func init() {
	f := BLhatse拉孜.(*拉孜LhatseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lhatse",
		TitleName: "拉孜",
		TitleCode: "b_lhatse",
	}
}
