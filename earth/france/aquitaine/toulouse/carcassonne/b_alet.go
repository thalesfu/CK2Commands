package carcassonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿莱AletBarony struct {
	feud.BaseBarony
}

var BAlet阿莱 feud.Barony = &阿莱AletBarony{}

func init() {
    f := BAlet阿莱.(*阿莱AletBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alet",
		TitleName: "阿莱",
		TitleCode: "b_alet",
	}
}
