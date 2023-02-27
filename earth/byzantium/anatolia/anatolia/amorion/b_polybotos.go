package amorion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波利波图斯PolybotosBarony struct {
	feud.BaseBarony
}

var BPolybotos波利波图斯 feud.Barony = &波利波图斯PolybotosBarony{}

func init() {
    f := BPolybotos波利波图斯.(*波利波图斯PolybotosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polybotos",
		TitleName: "波利波图斯",
		TitleCode: "b_polybotos",
	}
}
