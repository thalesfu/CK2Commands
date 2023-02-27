package amol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿赫斯萨克AkhsisakBarony struct {
	feud.BaseBarony
}

var BAkhsisak阿赫斯萨克 feud.Barony = &阿赫斯萨克AkhsisakBarony{}

func init() {
    f := BAkhsisak阿赫斯萨克.(*阿赫斯萨克AkhsisakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akhsisak",
		TitleName: "阿赫斯萨克",
		TitleCode: "b_akhsisak",
	}
}
