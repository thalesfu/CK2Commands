package makran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡里姆巴克什Karim_bakhshBarony struct {
	feud.BaseBarony
}

var BKarim_bakhsh卡里姆巴克什 feud.Barony = &卡里姆巴克什Karim_bakhshBarony{}

func init() {
    f := BKarim_bakhsh卡里姆巴克什.(*卡里姆巴克什Karim_bakhshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karim_bakhsh",
		TitleName: "卡里姆巴克什",
		TitleCode: "b_karim_bakhsh",
	}
}
