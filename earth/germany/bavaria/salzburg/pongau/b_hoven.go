package pongau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍文HovenBarony struct {
	feud.BaseBarony
}

var BHoven霍文 feud.Barony = &霍文HovenBarony{}

func init() {
	f := BHoven霍文.(*霍文HovenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hoven",
		TitleName: "霍文",
		TitleCode: "b_hoven",
	}
}
