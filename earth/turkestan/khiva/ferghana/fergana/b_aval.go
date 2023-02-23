package fergana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿瓦尔AvalBarony struct {
	feud.BaseBarony
}

var BAval阿瓦尔 feud.Barony = &阿瓦尔AvalBarony{}

func init() {
	f := BAval阿瓦尔.(*阿瓦尔AvalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aval",
		TitleName: "阿瓦尔",
		TitleCode: "b_aval",
	}
}
