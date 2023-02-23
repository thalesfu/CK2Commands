package palmyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉克ArakBarony struct {
	feud.BaseBarony
}

var BArak阿拉克 feud.Barony = &阿拉克ArakBarony{}

func init() {
	f := BArak阿拉克.(*阿拉克ArakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arak",
		TitleName: "阿拉克",
		TitleCode: "b_arak",
	}
}
