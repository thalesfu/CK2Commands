package apulia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢卡诺LucanoBarony struct {
	feud.BaseBarony
}

var BLucano卢卡诺 feud.Barony = &卢卡诺LucanoBarony{}

func init() {
    f := BLucano卢卡诺.(*卢卡诺LucanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lucano",
		TitleName: "卢卡诺",
		TitleCode: "b_lucano",
	}
}
