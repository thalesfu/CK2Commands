package krakowskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉科夫KrakowBarony struct {
	feud.BaseBarony
}

var BKrakow克拉科夫 feud.Barony = &克拉科夫KrakowBarony{}

func init() {
    f := BKrakow克拉科夫.(*克拉科夫KrakowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krakow",
		TitleName: "克拉科夫",
		TitleCode: "b_krakow",
	}
}
