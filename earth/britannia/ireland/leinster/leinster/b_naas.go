package leinster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内斯NaasBarony struct {
	feud.BaseBarony
}

var BNaas内斯 feud.Barony = &内斯NaasBarony{}

func init() {
    f := BNaas内斯.(*内斯NaasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naas",
		TitleName: "内斯",
		TitleCode: "b_naas",
	}
}
