package wiltshire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 威尔顿WiltonBarony struct {
	feud.BaseBarony
}

var BWilton威尔顿 feud.Barony = &威尔顿WiltonBarony{}

func init() {
	f := BWilton威尔顿.(*威尔顿WiltonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wilton",
		TitleName: "威尔顿",
		TitleCode: "b_wilton",
	}
}
