package simaramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 跋多那PattanaBarony struct {
	feud.BaseBarony
}

var BPattana跋多那 feud.Barony = &跋多那PattanaBarony{}

func init() {
	f := BPattana跋多那.(*跋多那PattanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pattana",
		TitleName: "跋多那",
		TitleCode: "b_pattana",
	}
}
