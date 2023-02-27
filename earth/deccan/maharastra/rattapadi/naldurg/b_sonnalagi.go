package naldurg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 十六城SonnalagiBarony struct {
	feud.BaseBarony
}

var BSonnalagi十六城 feud.Barony = &十六城SonnalagiBarony{}

func init() {
    f := BSonnalagi十六城.(*十六城SonnalagiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sonnalagi",
		TitleName: "十六城",
		TitleCode: "b_sonnalagi",
	}
}
