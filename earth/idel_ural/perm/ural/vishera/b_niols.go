package vishera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼奥尔斯NiolsBarony struct {
	feud.BaseBarony
}

var BNiols尼奥尔斯 feud.Barony = &尼奥尔斯NiolsBarony{}

func init() {
    f := BNiols尼奥尔斯.(*尼奥尔斯NiolsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niols",
		TitleName: "尼奥尔斯",
		TitleCode: "b_niols",
	}
}
