package negev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅因KmehinBarony struct {
	feud.BaseBarony
}

var BKmehin梅因 feud.Barony = &梅因KmehinBarony{}

func init() {
    f := BKmehin梅因.(*梅因KmehinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kmehin",
		TitleName: "梅因",
		TitleCode: "b_kmehin",
	}
}
