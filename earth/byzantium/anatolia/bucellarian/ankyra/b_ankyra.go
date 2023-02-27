package ankyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安居拉AnkyraBarony struct {
	feud.BaseBarony
}

var BAnkyra安居拉 feud.Barony = &安居拉AnkyraBarony{}

func init() {
    f := BAnkyra安居拉.(*安居拉AnkyraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ankyra",
		TitleName: "安居拉",
		TitleCode: "b_ankyra",
	}
}
