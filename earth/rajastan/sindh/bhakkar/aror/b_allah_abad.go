package aror

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉阿巴德Allah_abadBarony struct {
	feud.BaseBarony
}

var BAllah_abad阿拉阿巴德 feud.Barony = &阿拉阿巴德Allah_abadBarony{}

func init() {
    f := BAllah_abad阿拉阿巴德.(*阿拉阿巴德Allah_abadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "allah_abad",
		TitleName: "阿拉阿巴德",
		TitleCode: "b_allah_abad",
	}
}
