package chach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 白水城IsbijabBarony struct {
	feud.BaseBarony
}

var BIsbijab白水城 feud.Barony = &白水城IsbijabBarony{}

func init() {
    f := BIsbijab白水城.(*白水城IsbijabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isbijab",
		TitleName: "白水城",
		TitleCode: "b_isbijab",
	}
}
