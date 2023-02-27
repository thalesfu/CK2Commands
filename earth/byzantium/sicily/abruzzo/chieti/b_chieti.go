package chieti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基耶蒂ChietiBarony struct {
	feud.BaseBarony
}

var BChieti基耶蒂 feud.Barony = &基耶蒂ChietiBarony{}

func init() {
    f := BChieti基耶蒂.(*基耶蒂ChietiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chieti",
		TitleName: "基耶蒂",
		TitleCode: "b_chieti",
	}
}
