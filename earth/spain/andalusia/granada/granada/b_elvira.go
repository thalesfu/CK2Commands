package granada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔薇拉ElviraBarony struct {
	feud.BaseBarony
}

var BElvira埃尔薇拉 feud.Barony = &埃尔薇拉ElviraBarony{}

func init() {
    f := BElvira埃尔薇拉.(*埃尔薇拉ElviraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elvira",
		TitleName: "埃尔薇拉",
		TitleCode: "b_elvira",
	}
}
