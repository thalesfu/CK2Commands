package karashar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尉犁WeiliBarony struct {
	feud.BaseBarony
}

var BWeili尉犁 feud.Barony = &尉犁WeiliBarony{}

func init() {
    f := BWeili尉犁.(*尉犁WeiliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weili",
		TitleName: "尉犁",
		TitleCode: "b_weili",
	}
}
