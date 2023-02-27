package banavasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯灵盖里SringeriBarony struct {
	feud.BaseBarony
}

var BSringeri斯灵盖里 feud.Barony = &斯灵盖里SringeriBarony{}

func init() {
    f := BSringeri斯灵盖里.(*斯灵盖里SringeriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sringeri",
		TitleName: "斯灵盖里",
		TitleCode: "b_sringeri",
	}
}
