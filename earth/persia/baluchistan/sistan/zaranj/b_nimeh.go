package zaranj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼梅希NimehBarony struct {
	feud.BaseBarony
}

var BNimeh尼梅希 feud.Barony = &尼梅希NimehBarony{}

func init() {
	f := BNimeh尼梅希.(*尼梅希NimehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nimeh",
		TitleName: "尼梅希",
		TitleCode: "b_nimeh",
	}
}
