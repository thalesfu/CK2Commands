package arezzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢奇尼亚诺LucignanoBarony struct {
	feud.BaseBarony
}

var BLucignano卢奇尼亚诺 feud.Barony = &卢奇尼亚诺LucignanoBarony{}

func init() {
    f := BLucignano卢奇尼亚诺.(*卢奇尼亚诺LucignanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lucignano",
		TitleName: "卢奇尼亚诺",
		TitleCode: "b_lucignano",
	}
}
