package vidin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维舍斯拉夫ViseslavBarony struct {
	feud.BaseBarony
}

var BViseslav维舍斯拉夫 feud.Barony = &维舍斯拉夫ViseslavBarony{}

func init() {
    f := BViseslav维舍斯拉夫.(*维舍斯拉夫ViseslavBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viseslav",
		TitleName: "维舍斯拉夫",
		TitleCode: "b_viseslav",
	}
}
