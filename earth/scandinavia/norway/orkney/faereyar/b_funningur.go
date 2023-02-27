package faereyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富宁古尔FunningurBarony struct {
	feud.BaseBarony
}

var BFunningur富宁古尔 feud.Barony = &富宁古尔FunningurBarony{}

func init() {
    f := BFunningur富宁古尔.(*富宁古尔FunningurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "funningur",
		TitleName: "富宁古尔",
		TitleCode: "b_funningur",
	}
}
