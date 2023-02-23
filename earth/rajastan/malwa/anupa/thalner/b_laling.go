package thalner

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗尼恩LalingBarony struct {
	feud.BaseBarony
}

var BLaling罗尼恩 feud.Barony = &罗尼恩LalingBarony{}

func init() {
	f := BLaling罗尼恩.(*罗尼恩LalingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laling",
		TitleName: "罗尼恩",
		TitleCode: "b_laling",
	}
}
