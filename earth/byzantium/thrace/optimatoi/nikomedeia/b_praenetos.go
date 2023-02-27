package nikomedeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普莱PraenetosBarony struct {
	feud.BaseBarony
}

var BPraenetos普莱 feud.Barony = &普莱PraenetosBarony{}

func init() {
    f := BPraenetos普莱.(*普莱PraenetosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "praenetos",
		TitleName: "普莱",
		TitleCode: "b_praenetos",
	}
}
