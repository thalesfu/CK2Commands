package sakya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 龙中LungrongBarony struct {
	feud.BaseBarony
}

var BLungrong龙中 feud.Barony = &龙中LungrongBarony{}

func init() {
    f := BLungrong龙中.(*龙中LungrongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lungrong",
		TitleName: "龙中",
		TitleCode: "b_lungrong",
	}
}
