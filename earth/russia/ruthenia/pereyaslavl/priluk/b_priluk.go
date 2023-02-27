package priluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普里卢基PrilukBarony struct {
	feud.BaseBarony
}

var BPriluk普里卢基 feud.Barony = &普里卢基PrilukBarony{}

func init() {
    f := BPriluk普里卢基.(*普里卢基PrilukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "priluk",
		TitleName: "普里卢基",
		TitleCode: "b_priluk",
	}
}
