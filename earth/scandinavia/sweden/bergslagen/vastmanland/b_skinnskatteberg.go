package vastmanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欣斯卡特贝里SkinnskattebergBarony struct {
	feud.BaseBarony
}

var BSkinnskatteberg欣斯卡特贝里 feud.Barony = &欣斯卡特贝里SkinnskattebergBarony{}

func init() {
	f := BSkinnskatteberg欣斯卡特贝里.(*欣斯卡特贝里SkinnskattebergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skinnskatteberg",
		TitleName: "欣斯卡特贝里",
		TitleCode: "b_skinnskatteberg",
	}
}
