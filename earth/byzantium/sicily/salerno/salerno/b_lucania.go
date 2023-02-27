package salerno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢卡尼亚LucaniaBarony struct {
	feud.BaseBarony
}

var BLucania卢卡尼亚 feud.Barony = &卢卡尼亚LucaniaBarony{}

func init() {
    f := BLucania卢卡尼亚.(*卢卡尼亚LucaniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lucania",
		TitleName: "卢卡尼亚",
		TitleCode: "b_lucania",
	}
}
