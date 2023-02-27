package fachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古洛GouloBarony struct {
	feud.BaseBarony
}

var BGoulo古洛 feud.Barony = &古洛GouloBarony{}

func init() {
    f := BGoulo古洛.(*古洛GouloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goulo",
		TitleName: "古洛",
		TitleCode: "b_goulo",
	}
}
