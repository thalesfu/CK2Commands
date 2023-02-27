package uwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡特拉SitraBarony struct {
	feud.BaseBarony
}

var BSitra锡特拉 feud.Barony = &锡特拉SitraBarony{}

func init() {
    f := BSitra锡特拉.(*锡特拉SitraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sitra",
		TitleName: "锡特拉",
		TitleCode: "b_sitra",
	}
}
