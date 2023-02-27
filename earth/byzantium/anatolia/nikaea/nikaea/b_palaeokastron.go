package nikaea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普莱欧卡斯特朗PalaeokastronBarony struct {
	feud.BaseBarony
}

var BPalaeokastron普莱欧卡斯特朗 feud.Barony = &普莱欧卡斯特朗PalaeokastronBarony{}

func init() {
    f := BPalaeokastron普莱欧卡斯特朗.(*普莱欧卡斯特朗PalaeokastronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palaeokastron",
		TitleName: "普莱欧卡斯特朗",
		TitleCode: "b_palaeokastron",
	}
}
