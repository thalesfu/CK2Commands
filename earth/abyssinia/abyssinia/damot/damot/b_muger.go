package damot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆格尔MugerBarony struct {
	feud.BaseBarony
}

var BMuger穆格尔 feud.Barony = &穆格尔MugerBarony{}

func init() {
    f := BMuger穆格尔.(*穆格尔MugerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muger",
		TitleName: "穆格尔",
		TitleCode: "b_muger",
	}
}
