package turkestan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉尔库姆AralkumBarony struct {
	feud.BaseBarony
}

var BAralkum阿拉尔库姆 feud.Barony = &阿拉尔库姆AralkumBarony{}

func init() {
	f := BAralkum阿拉尔库姆.(*阿拉尔库姆AralkumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aralkum",
		TitleName: "阿拉尔库姆",
		TitleCode: "b_aralkum",
	}
}
