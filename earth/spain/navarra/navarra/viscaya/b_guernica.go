package viscaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格尔尼卡GuernicaBarony struct {
	feud.BaseBarony
}

var BGuernica格尔尼卡 feud.Barony = &格尔尼卡GuernicaBarony{}

func init() {
	f := BGuernica格尔尼卡.(*格尔尼卡GuernicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guernica",
		TitleName: "格尔尼卡",
		TitleCode: "b_guernica",
	}
}
