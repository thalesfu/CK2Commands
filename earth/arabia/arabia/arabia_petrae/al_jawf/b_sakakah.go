package al_jawf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞卡凯SakakahBarony struct {
	feud.BaseBarony
}

var BSakakah塞卡凯 feud.Barony = &塞卡凯SakakahBarony{}

func init() {
    f := BSakakah塞卡凯.(*塞卡凯SakakahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sakakah",
		TitleName: "塞卡凯",
		TitleCode: "b_sakakah",
	}
}
