package orekhovo_zouievo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图马TumaBarony struct {
	feud.BaseBarony
}

var BTuma图马 feud.Barony = &图马TumaBarony{}

func init() {
    f := BTuma图马.(*图马TumaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuma",
		TitleName: "图马",
		TitleCode: "b_tuma",
	}
}
