package lower_volga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡梅申KamyshinBarony struct {
	feud.BaseBarony
}

var BKamyshin卡梅申 feud.Barony = &卡梅申KamyshinBarony{}

func init() {
    f := BKamyshin卡梅申.(*卡梅申KamyshinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamyshin",
		TitleName: "卡梅申",
		TitleCode: "b_kamyshin",
	}
}
