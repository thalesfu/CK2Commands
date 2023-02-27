package oromieh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦吉克AvajiqBarony struct {
	feud.BaseBarony
}

var BAvajiq阿瓦吉克 feud.Barony = &阿瓦吉克AvajiqBarony{}

func init() {
    f := BAvajiq阿瓦吉克.(*阿瓦吉克AvajiqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avajiq",
		TitleName: "阿瓦吉克",
		TitleCode: "b_avajiq",
	}
}
