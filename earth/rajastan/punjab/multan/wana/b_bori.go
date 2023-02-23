package wana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩里BoriBarony struct {
	feud.BaseBarony
}

var BBori菩里 feud.Barony = &菩里BoriBarony{}

func init() {
	f := BBori菩里.(*菩里BoriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bori",
		TitleName: "菩里",
		TitleCode: "b_bori",
	}
}
