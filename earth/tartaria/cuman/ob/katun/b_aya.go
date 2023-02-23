package katun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿亚AyaBarony struct {
	feud.BaseBarony
}

var BAya阿亚 feud.Barony = &阿亚AyaBarony{}

func init() {
	f := BAya阿亚.(*阿亚AyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aya",
		TitleName: "阿亚",
		TitleCode: "b_aya",
	}
}
