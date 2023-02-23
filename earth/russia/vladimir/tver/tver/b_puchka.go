package tver

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普奇卡PuchkaBarony struct {
	feud.BaseBarony
}

var BPuchka普奇卡 feud.Barony = &普奇卡PuchkaBarony{}

func init() {
	f := BPuchka普奇卡.(*普奇卡PuchkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puchka",
		TitleName: "普奇卡",
		TitleCode: "b_puchka",
	}
}
