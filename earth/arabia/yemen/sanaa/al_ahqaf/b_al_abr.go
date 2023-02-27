package al_ahqaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卜尔堡Al_abrBarony struct {
	feud.BaseBarony
}

var BAl_abr阿卜尔堡 feud.Barony = &阿卜尔堡Al_abrBarony{}

func init() {
    f := BAl_abr阿卜尔堡.(*阿卜尔堡Al_abrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_abr",
		TitleName: "阿卜尔堡",
		TitleCode: "b_al_abr",
	}
}
