package duqm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班克BankBarony struct {
	feud.BaseBarony
}

var BBank班克 feud.Barony = &班克BankBarony{}

func init() {
    f := BBank班克.(*班克BankBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bank",
		TitleName: "班克",
		TitleCode: "b_bank",
	}
}
