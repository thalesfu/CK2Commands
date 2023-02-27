package salzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳芬LaufenBarony struct {
	feud.BaseBarony
}

var BLaufen劳芬 feud.Barony = &劳芬LaufenBarony{}

func init() {
    f := BLaufen劳芬.(*劳芬LaufenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laufen",
		TitleName: "劳芬",
		TitleCode: "b_laufen",
	}
}
