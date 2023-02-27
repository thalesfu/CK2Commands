package canda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉久拉RajuraBarony struct {
	feud.BaseBarony
}

var BRajura拉久拉 feud.Barony = &拉久拉RajuraBarony{}

func init() {
    f := BRajura拉久拉.(*拉久拉RajuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajura",
		TitleName: "拉久拉",
		TitleCode: "b_rajura",
	}
}
