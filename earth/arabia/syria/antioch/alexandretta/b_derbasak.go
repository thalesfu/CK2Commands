package alexandretta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔伯萨克DerbasakBarony struct {
	feud.BaseBarony
}

var BDerbasak达尔伯萨克 feud.Barony = &达尔伯萨克DerbasakBarony{}

func init() {
	f := BDerbasak达尔伯萨克.(*达尔伯萨克DerbasakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "derbasak",
		TitleName: "达尔伯萨克",
		TitleCode: "b_derbasak",
	}
}
