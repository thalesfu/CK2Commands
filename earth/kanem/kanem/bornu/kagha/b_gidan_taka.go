package kagha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉丹塔卡Gidan_takaBarony struct {
	feud.BaseBarony
}

var BGidan_taka吉丹塔卡 feud.Barony = &吉丹塔卡Gidan_takaBarony{}

func init() {
    f := BGidan_taka吉丹塔卡.(*吉丹塔卡Gidan_takaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gidan_taka",
		TitleName: "吉丹塔卡",
		TitleCode: "b_gidan_taka",
	}
}
