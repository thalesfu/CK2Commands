package terek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马贾尔MajarBarony struct {
	feud.BaseBarony
}

var BMajar马贾尔 feud.Barony = &马贾尔MajarBarony{}

func init() {
    f := BMajar马贾尔.(*马贾尔MajarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "majar",
		TitleName: "马贾尔",
		TitleCode: "b_majar",
	}
}
