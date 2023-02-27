package mtskheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哥里GoriBarony struct {
	feud.BaseBarony
}

var BGori哥里 feud.Barony = &哥里GoriBarony{}

func init() {
    f := BGori哥里.(*哥里GoriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gori",
		TitleName: "哥里",
		TitleCode: "b_gori",
	}
}
