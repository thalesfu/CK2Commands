package kremenchuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥塔尔卡OtarkaBarony struct {
	feud.BaseBarony
}

var BOtarka奥塔尔卡 feud.Barony = &奥塔尔卡OtarkaBarony{}

func init() {
    f := BOtarka奥塔尔卡.(*奥塔尔卡OtarkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "otarka",
		TitleName: "奥塔尔卡",
		TitleCode: "b_otarka",
	}
}
