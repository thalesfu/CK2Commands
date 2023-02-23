package onega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥努斯AunusBarony struct {
	feud.BaseBarony
}

var BAunus奥努斯 feud.Barony = &奥努斯AunusBarony{}

func init() {
	f := BAunus奥努斯.(*奥努斯AunusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aunus",
		TitleName: "奥努斯",
		TitleCode: "b_aunus",
	}
}
