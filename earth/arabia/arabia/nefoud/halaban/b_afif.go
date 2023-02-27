package halaban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿菲夫AfifBarony struct {
	feud.BaseBarony
}

var BAfif阿菲夫 feud.Barony = &阿菲夫AfifBarony{}

func init() {
    f := BAfif阿菲夫.(*阿菲夫AfifBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "afif",
		TitleName: "阿菲夫",
		TitleCode: "b_afif",
	}
}
