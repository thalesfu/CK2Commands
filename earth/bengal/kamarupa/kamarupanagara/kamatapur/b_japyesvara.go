package kamatapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇比曳湿伐罗JapyesvaraBarony struct {
	feud.BaseBarony
}

var BJapyesvara阇比曳湿伐罗 feud.Barony = &阇比曳湿伐罗JapyesvaraBarony{}

func init() {
    f := BJapyesvara阇比曳湿伐罗.(*阇比曳湿伐罗JapyesvaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "japyesvara",
		TitleName: "阇比曳湿伐罗",
		TitleCode: "b_japyesvara",
	}
}
