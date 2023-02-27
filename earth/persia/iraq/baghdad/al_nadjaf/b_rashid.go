package al_nadjaf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗塞塔RashidBarony struct {
	feud.BaseBarony
}

var BRashid罗塞塔 feud.Barony = &罗塞塔RashidBarony{}

func init() {
    f := BRashid罗塞塔.(*罗塞塔RashidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rashid",
		TitleName: "罗塞塔",
		TitleCode: "b_rashid",
	}
}
