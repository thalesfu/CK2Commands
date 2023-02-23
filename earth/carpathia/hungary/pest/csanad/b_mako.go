package csanad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毛科MakoBarony struct {
	feud.BaseBarony
}

var BMako毛科 feud.Barony = &毛科MakoBarony{}

func init() {
	f := BMako毛科.(*毛科MakoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mako",
		TitleName: "毛科",
		TitleCode: "b_mako",
	}
}
