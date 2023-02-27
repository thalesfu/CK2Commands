package benevento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿韦利诺AvellinoBarony struct {
	feud.BaseBarony
}

var BAvellino阿韦利诺 feud.Barony = &阿韦利诺AvellinoBarony{}

func init() {
    f := BAvellino阿韦利诺.(*阿韦利诺AvellinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avellino",
		TitleName: "阿韦利诺",
		TitleCode: "b_avellino",
	}
}
