package paphlagonia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 匝利刻斯ZaliscusBarony struct {
	feud.BaseBarony
}

var BZaliscus匝利刻斯 feud.Barony = &匝利刻斯ZaliscusBarony{}

func init() {
	f := BZaliscus匝利刻斯.(*匝利刻斯ZaliscusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaliscus",
		TitleName: "匝利刻斯",
		TitleCode: "b_zaliscus",
	}
}
