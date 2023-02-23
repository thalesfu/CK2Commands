package rahbah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼梅赫NimrehBarony struct {
	feud.BaseBarony
}

var BNimreh尼梅赫 feud.Barony = &尼梅赫NimrehBarony{}

func init() {
	f := BNimreh尼梅赫.(*尼梅赫NimrehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nimreh",
		TitleName: "尼梅赫",
		TitleCode: "b_nimreh",
	}
}
