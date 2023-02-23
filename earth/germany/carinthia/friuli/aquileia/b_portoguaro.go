package aquileia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔托格鲁阿罗PortoguaroBarony struct {
	feud.BaseBarony
}

var BPortoguaro波尔托格鲁阿罗 feud.Barony = &波尔托格鲁阿罗PortoguaroBarony{}

func init() {
	f := BPortoguaro波尔托格鲁阿罗.(*波尔托格鲁阿罗PortoguaroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "portoguaro",
		TitleName: "波尔托格鲁阿罗",
		TitleCode: "b_portoguaro",
	}
}
