package wight

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纽波特Newport_wightBarony struct {
	feud.BaseBarony
}

var BNewport_wight纽波特 feud.Barony = &纽波特Newport_wightBarony{}

func init() {
    f := BNewport_wight纽波特.(*纽波特Newport_wightBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "newport_wight",
		TitleName: "纽波特",
		TitleCode: "b_newport_wight",
	}
}
