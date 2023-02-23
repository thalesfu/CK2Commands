package serdica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普拉韦茨PravetsBarony struct {
	feud.BaseBarony
}

var BPravets普拉韦茨 feud.Barony = &普拉韦茨PravetsBarony{}

func init() {
	f := BPravets普拉韦茨.(*普拉韦茨PravetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pravets",
		TitleName: "普拉韦茨",
		TitleCode: "b_pravets",
	}
}
