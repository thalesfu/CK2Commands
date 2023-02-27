package cinarca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普罗普里亚诺ProprianoBarony struct {
	feud.BaseBarony
}

var BPropriano普罗普里亚诺 feud.Barony = &普罗普里亚诺ProprianoBarony{}

func init() {
    f := BPropriano普罗普里亚诺.(*普罗普里亚诺ProprianoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "propriano",
		TitleName: "普罗普里亚诺",
		TitleCode: "b_propriano",
	}
}
