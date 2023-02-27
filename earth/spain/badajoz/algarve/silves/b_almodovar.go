package silves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔莫多瓦AlmodovarBarony struct {
	feud.BaseBarony
}

var BAlmodovar阿尔莫多瓦 feud.Barony = &阿尔莫多瓦AlmodovarBarony{}

func init() {
    f := BAlmodovar阿尔莫多瓦.(*阿尔莫多瓦AlmodovarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almodovar",
		TitleName: "阿尔莫多瓦",
		TitleCode: "b_almodovar",
	}
}
