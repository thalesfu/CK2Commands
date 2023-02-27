package cieszyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普什奇纳PszczynaBarony struct {
	feud.BaseBarony
}

var BPszczyna普什奇纳 feud.Barony = &普什奇纳PszczynaBarony{}

func init() {
    f := BPszczyna普什奇纳.(*普什奇纳PszczynaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pszczyna",
		TitleName: "普什奇纳",
		TitleCode: "b_pszczyna",
	}
}
