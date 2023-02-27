package ulm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比伯拉赫BiberachBarony struct {
	feud.BaseBarony
}

var BBiberach比伯拉赫 feud.Barony = &比伯拉赫BiberachBarony{}

func init() {
    f := BBiberach比伯拉赫.(*比伯拉赫BiberachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biberach",
		TitleName: "比伯拉赫",
		TitleCode: "b_biberach",
	}
}
