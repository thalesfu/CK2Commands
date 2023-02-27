package dulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 热水ReshuiBarony struct {
	feud.BaseBarony
}

var BReshui热水 feud.Barony = &热水ReshuiBarony{}

func init() {
    f := BReshui热水.(*热水ReshuiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reshui",
		TitleName: "热水",
		TitleCode: "b_reshui",
	}
}
