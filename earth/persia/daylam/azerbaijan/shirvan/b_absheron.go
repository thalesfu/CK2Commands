package shirvan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿普舍龙AbsheronBarony struct {
	feud.BaseBarony
}

var BAbsheron阿普舍龙 feud.Barony = &阿普舍龙AbsheronBarony{}

func init() {
    f := BAbsheron阿普舍龙.(*阿普舍龙AbsheronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "absheron",
		TitleName: "阿普舍龙",
		TitleCode: "b_absheron",
	}
}
