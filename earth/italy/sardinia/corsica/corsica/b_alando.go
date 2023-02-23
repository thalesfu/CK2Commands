package corsica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿兰多AlandoBarony struct {
	feud.BaseBarony
}

var BAlando阿兰多 feud.Barony = &阿兰多AlandoBarony{}

func init() {
	f := BAlando阿兰多.(*阿兰多AlandoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alando",
		TitleName: "阿兰多",
		TitleCode: "b_alando",
	}
}
