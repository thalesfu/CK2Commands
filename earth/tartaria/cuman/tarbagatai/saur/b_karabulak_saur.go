package saur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉布拉克Karabulak_saurBarony struct {
	feud.BaseBarony
}

var BKarabulak_saur卡拉布拉克 feud.Barony = &卡拉布拉克Karabulak_saurBarony{}

func init() {
    f := BKarabulak_saur卡拉布拉克.(*卡拉布拉克Karabulak_saurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karabulak_saur",
		TitleName: "卡拉布拉克",
		TitleCode: "b_karabulak_saur",
	}
}
