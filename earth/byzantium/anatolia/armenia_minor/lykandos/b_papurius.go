package lykandos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉普乌斯PapuriusBarony struct {
	feud.BaseBarony
}

var BPapurius拉普乌斯 feud.Barony = &拉普乌斯PapuriusBarony{}

func init() {
    f := BPapurius拉普乌斯.(*拉普乌斯PapuriusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "papurius",
		TitleName: "拉普乌斯",
		TitleCode: "b_papurius",
	}
}
