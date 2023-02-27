package kremenchuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡特梅什KatmyshBarony struct {
	feud.BaseBarony
}

var BKatmysh卡特梅什 feud.Barony = &卡特梅什KatmyshBarony{}

func init() {
    f := BKatmysh卡特梅什.(*卡特梅什KatmyshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katmysh",
		TitleName: "卡特梅什",
		TitleCode: "b_katmysh",
	}
}
