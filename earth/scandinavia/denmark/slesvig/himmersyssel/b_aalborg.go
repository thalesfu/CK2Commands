package himmersyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔堡AalborgBarony struct {
	feud.BaseBarony
}

var BAalborg奥尔堡 feud.Barony = &奥尔堡AalborgBarony{}

func init() {
	f := BAalborg奥尔堡.(*奥尔堡AalborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aalborg",
		TitleName: "奥尔堡",
		TitleCode: "b_aalborg",
	}
}
