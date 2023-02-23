package srirangapatna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室罗伐拿白泽ShravanabelagolaBarony struct {
	feud.BaseBarony
}

var BShravanabelagola室罗伐拿白泽 feud.Barony = &室罗伐拿白泽ShravanabelagolaBarony{}

func init() {
	f := BShravanabelagola室罗伐拿白泽.(*室罗伐拿白泽ShravanabelagolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shravanabelagola",
		TitleName: "室罗伐拿白泽",
		TitleCode: "b_shravanabelagola",
	}
}
