package irbid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加大拉GaderaBarony struct {
	feud.BaseBarony
}

var BGadera加大拉 feud.Barony = &加大拉GaderaBarony{}

func init() {
	f := BGadera加大拉.(*加大拉GaderaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gadera",
		TitleName: "加大拉",
		TitleCode: "b_gadera",
	}
}
