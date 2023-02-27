package lori

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛里堡Lori_berdBarony struct {
	feud.BaseBarony
}

var BLori_berd洛里堡 feud.Barony = &洛里堡Lori_berdBarony{}

func init() {
    f := BLori_berd洛里堡.(*洛里堡Lori_berdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lori_berd",
		TitleName: "洛里堡",
		TitleCode: "b_lori_berd",
	}
}
