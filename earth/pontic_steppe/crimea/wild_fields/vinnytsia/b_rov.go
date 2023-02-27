package vinnytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗夫RovBarony struct {
	feud.BaseBarony
}

var BRov罗夫 feud.Barony = &罗夫RovBarony{}

func init() {
    f := BRov罗夫.(*罗夫RovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rov",
		TitleName: "罗夫",
		TitleCode: "b_rov",
	}
}
