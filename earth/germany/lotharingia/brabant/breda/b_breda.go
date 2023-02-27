package breda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布雷达BredaBarony struct {
	feud.BaseBarony
}

var BBreda布雷达 feud.Barony = &布雷达BredaBarony{}

func init() {
    f := BBreda布雷达.(*布雷达BredaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "breda",
		TitleName: "布雷达",
		TitleCode: "b_breda",
	}
}
