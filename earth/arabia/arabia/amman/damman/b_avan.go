package damman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦恩AvanBarony struct {
	feud.BaseBarony
}

var BAvan阿瓦恩 feud.Barony = &阿瓦恩AvanBarony{}

func init() {
	f := BAvan阿瓦恩.(*阿瓦恩AvanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avan",
		TitleName: "阿瓦恩",
		TitleCode: "b_avan",
	}
}
