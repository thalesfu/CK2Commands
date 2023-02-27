package nyima

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼玛NyimaBarony struct {
	feud.BaseBarony
}

var BNyima尼玛 feud.Barony = &尼玛NyimaBarony{}

func init() {
    f := BNyima尼玛.(*尼玛NyimaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyima",
		TitleName: "尼玛",
		TitleCode: "b_nyima",
	}
}
