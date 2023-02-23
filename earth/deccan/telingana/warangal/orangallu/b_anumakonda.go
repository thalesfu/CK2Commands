package orangallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿奴摩军荼AnumakondaBarony struct {
	feud.BaseBarony
}

var BAnumakonda阿奴摩军荼 feud.Barony = &阿奴摩军荼AnumakondaBarony{}

func init() {
	f := BAnumakonda阿奴摩军荼.(*阿奴摩军荼AnumakondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anumakonda",
		TitleName: "阿奴摩军荼",
		TitleCode: "b_anumakonda",
	}
}
