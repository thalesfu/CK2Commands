package romerike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费聪FetsundBarony struct {
	feud.BaseBarony
}

var BFetsund费聪 feud.Barony = &费聪FetsundBarony{}

func init() {
    f := BFetsund费聪.(*费聪FetsundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fetsund",
		TitleName: "费聪",
		TitleCode: "b_fetsund",
	}
}
