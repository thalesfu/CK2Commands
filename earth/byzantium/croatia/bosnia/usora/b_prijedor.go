package usora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普里耶多尔PrijedorBarony struct {
	feud.BaseBarony
}

var BPrijedor普里耶多尔 feud.Barony = &普里耶多尔PrijedorBarony{}

func init() {
    f := BPrijedor普里耶多尔.(*普里耶多尔PrijedorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prijedor",
		TitleName: "普里耶多尔",
		TitleCode: "b_prijedor",
	}
}
