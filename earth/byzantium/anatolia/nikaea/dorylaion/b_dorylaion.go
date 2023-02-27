package dorylaion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多律莱翁DorylaionBarony struct {
	feud.BaseBarony
}

var BDorylaion多律莱翁 feud.Barony = &多律莱翁DorylaionBarony{}

func init() {
    f := BDorylaion多律莱翁.(*多律莱翁DorylaionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dorylaion",
		TitleName: "多律莱翁",
		TitleCode: "b_dorylaion",
	}
}
