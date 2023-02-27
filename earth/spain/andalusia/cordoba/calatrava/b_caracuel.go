package calatrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉库埃尔CaracuelBarony struct {
	feud.BaseBarony
}

var BCaracuel卡拉库埃尔 feud.Barony = &卡拉库埃尔CaracuelBarony{}

func init() {
    f := BCaracuel卡拉库埃尔.(*卡拉库埃尔CaracuelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caracuel",
		TitleName: "卡拉库埃尔",
		TitleCode: "b_caracuel",
	}
}
