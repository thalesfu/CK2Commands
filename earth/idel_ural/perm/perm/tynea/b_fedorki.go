package tynea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费多尔基FedorkiBarony struct {
	feud.BaseBarony
}

var BFedorki费多尔基 feud.Barony = &费多尔基FedorkiBarony{}

func init() {
    f := BFedorki费多尔基.(*费多尔基FedorkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fedorki",
		TitleName: "费多尔基",
		TitleCode: "b_fedorki",
	}
}
