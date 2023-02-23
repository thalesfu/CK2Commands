package kirkuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈赫穆尔MakhmurBarony struct {
	feud.BaseBarony
}

var BMakhmur迈赫穆尔 feud.Barony = &迈赫穆尔MakhmurBarony{}

func init() {
	f := BMakhmur迈赫穆尔.(*迈赫穆尔MakhmurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makhmur",
		TitleName: "迈赫穆尔",
		TitleCode: "b_makhmur",
	}
}
