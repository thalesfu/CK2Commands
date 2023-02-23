package kota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆那VanaBarony struct {
	feud.BaseBarony
}

var BVana婆那 feud.Barony = &婆那VanaBarony{}

func init() {
	f := BVana婆那.(*婆那VanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vana",
		TitleName: "婆那",
		TitleCode: "b_vana",
	}
}
