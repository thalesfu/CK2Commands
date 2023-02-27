package annaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布乌尼ElbouniBarony struct {
	feud.BaseBarony
}

var BElbouni布乌尼 feud.Barony = &布乌尼ElbouniBarony{}

func init() {
    f := BElbouni布乌尼.(*布乌尼ElbouniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elbouni",
		TitleName: "布乌尼",
		TitleCode: "b_elbouni",
	}
}
