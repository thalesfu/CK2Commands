package badajoz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅里达MeridaBarony struct {
	feud.BaseBarony
}

var BMerida梅里达 feud.Barony = &梅里达MeridaBarony{}

func init() {
    f := BMerida梅里达.(*梅里达MeridaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merida",
		TitleName: "梅里达",
		TitleCode: "b_merida",
	}
}
