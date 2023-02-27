package miass

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉斯韦特RassvetBarony struct {
	feud.BaseBarony
}

var BRassvet拉斯韦特 feud.Barony = &拉斯韦特RassvetBarony{}

func init() {
    f := BRassvet拉斯韦特.(*拉斯韦特RassvetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rassvet",
		TitleName: "拉斯韦特",
		TitleCode: "b_rassvet",
	}
}
