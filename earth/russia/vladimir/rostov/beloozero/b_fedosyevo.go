package beloozero

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费多谢沃FedosyevoBarony struct {
	feud.BaseBarony
}

var BFedosyevo费多谢沃 feud.Barony = &费多谢沃FedosyevoBarony{}

func init() {
    f := BFedosyevo费多谢沃.(*费多谢沃FedosyevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fedosyevo",
		TitleName: "费多谢沃",
		TitleCode: "b_fedosyevo",
	}
}
