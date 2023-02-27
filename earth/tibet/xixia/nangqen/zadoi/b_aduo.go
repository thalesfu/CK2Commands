package zadoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿多AduoBarony struct {
	feud.BaseBarony
}

var BAduo阿多 feud.Barony = &阿多AduoBarony{}

func init() {
    f := BAduo阿多.(*阿多AduoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aduo",
		TitleName: "阿多",
		TitleCode: "b_aduo",
	}
}
