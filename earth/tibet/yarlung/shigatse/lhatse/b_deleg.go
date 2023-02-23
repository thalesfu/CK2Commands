package lhatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德来DelegBarony struct {
	feud.BaseBarony
}

var BDeleg德来 feud.Barony = &德来DelegBarony{}

func init() {
	f := BDeleg德来.(*德来DelegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deleg",
		TitleName: "德来",
		TitleCode: "b_deleg",
	}
}
