package zyriane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波萨德PosadBarony struct {
	feud.BaseBarony
}

var BPosad波萨德 feud.Barony = &波萨德PosadBarony{}

func init() {
    f := BPosad波萨德.(*波萨德PosadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "posad",
		TitleName: "波萨德",
		TitleCode: "b_posad",
	}
}
