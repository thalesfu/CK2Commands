package hanyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯迪尔AlkasdirBarony struct {
	feud.BaseBarony
}

var BAlkasdir卡斯迪尔 feud.Barony = &卡斯迪尔AlkasdirBarony{}

func init() {
	f := BAlkasdir卡斯迪尔.(*卡斯迪尔AlkasdirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alkasdir",
		TitleName: "卡斯迪尔",
		TitleCode: "b_alkasdir",
	}
}
