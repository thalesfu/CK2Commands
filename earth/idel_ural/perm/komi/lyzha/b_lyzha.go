package lyzha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷扎LyzhaBarony struct {
	feud.BaseBarony
}

var BLyzha雷扎 feud.Barony = &雷扎LyzhaBarony{}

func init() {
    f := BLyzha雷扎.(*雷扎LyzhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyzha",
		TitleName: "雷扎",
		TitleCode: "b_lyzha",
	}
}
