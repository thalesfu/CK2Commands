package badajoz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴达霍斯BadajozBarony struct {
	feud.BaseBarony
}

var BBadajoz巴达霍斯 feud.Barony = &巴达霍斯BadajozBarony{}

func init() {
    f := BBadajoz巴达霍斯.(*巴达霍斯BadajozBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "badajoz",
		TitleName: "巴达霍斯",
		TitleCode: "b_badajoz",
	}
}
