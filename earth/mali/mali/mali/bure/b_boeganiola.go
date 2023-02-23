package bure

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博埃加尼奥拉BoeganiolaBarony struct {
	feud.BaseBarony
}

var BBoeganiola博埃加尼奥拉 feud.Barony = &博埃加尼奥拉BoeganiolaBarony{}

func init() {
	f := BBoeganiola博埃加尼奥拉.(*博埃加尼奥拉BoeganiolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boeganiola",
		TitleName: "博埃加尼奥拉",
		TitleCode: "b_boeganiola",
	}
}
