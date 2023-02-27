package shemakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿内赫AnigBarony struct {
	feud.BaseBarony
}

var BAnig阿内赫 feud.Barony = &阿内赫AnigBarony{}

func init() {
    f := BAnig阿内赫.(*阿内赫AnigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anig",
		TitleName: "阿内赫",
		TitleCode: "b_anig",
	}
}
