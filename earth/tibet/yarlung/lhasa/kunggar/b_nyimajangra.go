package kunggar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼玛江热NyimajangraBarony struct {
	feud.BaseBarony
}

var BNyimajangra尼玛江热 feud.Barony = &尼玛江热NyimajangraBarony{}

func init() {
    f := BNyimajangra尼玛江热.(*尼玛江热NyimajangraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyimajangra",
		TitleName: "尼玛江热",
		TitleCode: "b_nyimajangra",
	}
}
