package kunlun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿其克AshiBarony struct {
	feud.BaseBarony
}

var BAshi阿其克 feud.Barony = &阿其克AshiBarony{}

func init() {
    f := BAshi阿其克.(*阿其克AshiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ashi",
		TitleName: "阿其克",
		TitleCode: "b_ashi",
	}
}
