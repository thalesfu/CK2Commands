package al_ahqaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 不阿詹BujanBarony struct {
	feud.BaseBarony
}

var BBujan不阿詹 feud.Barony = &不阿詹BujanBarony{}

func init() {
    f := BBujan不阿詹.(*不阿詹BujanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bujan",
		TitleName: "不阿詹",
		TitleCode: "b_bujan",
	}
}
