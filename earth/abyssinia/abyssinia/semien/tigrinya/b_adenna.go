package tigrinya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿丹那AdennaBarony struct {
	feud.BaseBarony
}

var BAdenna阿丹那 feud.Barony = &阿丹那AdennaBarony{}

func init() {
	f := BAdenna阿丹那.(*阿丹那AdennaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adenna",
		TitleName: "阿丹那",
		TitleCode: "b_adenna",
	}
}
