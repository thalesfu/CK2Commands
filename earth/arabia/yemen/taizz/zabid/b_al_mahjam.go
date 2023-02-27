package zabid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马赫贾姆Al_mahjamBarony struct {
	feud.BaseBarony
}

var BAl_mahjam马赫贾姆 feud.Barony = &马赫贾姆Al_mahjamBarony{}

func init() {
    f := BAl_mahjam马赫贾姆.(*马赫贾姆Al_mahjamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_mahjam",
		TitleName: "马赫贾姆",
		TitleCode: "b_al_mahjam",
	}
}
