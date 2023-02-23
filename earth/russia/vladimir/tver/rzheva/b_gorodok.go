package rzheva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈罗多克GorodokBarony struct {
	feud.BaseBarony
}

var BGorodok戈罗多克 feud.Barony = &戈罗多克GorodokBarony{}

func init() {
	f := BGorodok戈罗多克.(*戈罗多克GorodokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorodok",
		TitleName: "戈罗多克",
		TitleCode: "b_gorodok",
	}
}
