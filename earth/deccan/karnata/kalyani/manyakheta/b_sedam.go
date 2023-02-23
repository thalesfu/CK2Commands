package manyakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞达姆SedamBarony struct {
	feud.BaseBarony
}

var BSedam塞达姆 feud.Barony = &塞达姆SedamBarony{}

func init() {
	f := BSedam塞达姆.(*塞达姆SedamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sedam",
		TitleName: "塞达姆",
		TitleCode: "b_sedam",
	}
}
