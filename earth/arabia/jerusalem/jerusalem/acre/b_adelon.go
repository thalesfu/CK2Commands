package acre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿德朗AdelonBarony struct {
	feud.BaseBarony
}

var BAdelon阿德朗 feud.Barony = &阿德朗AdelonBarony{}

func init() {
    f := BAdelon阿德朗.(*阿德朗AdelonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adelon",
		TitleName: "阿德朗",
		TitleCode: "b_adelon",
	}
}
