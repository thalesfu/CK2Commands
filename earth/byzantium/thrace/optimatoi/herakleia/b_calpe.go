package herakleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔佩CalpeBarony struct {
	feud.BaseBarony
}

var BCalpe卡尔佩 feud.Barony = &卡尔佩CalpeBarony{}

func init() {
	f := BCalpe卡尔佩.(*卡尔佩CalpeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calpe",
		TitleName: "卡尔佩",
		TitleCode: "b_calpe",
	}
}
