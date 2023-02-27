package dariya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德拉伊耶DariyaBarony struct {
	feud.BaseBarony
}

var BDariya德拉伊耶 feud.Barony = &德拉伊耶DariyaBarony{}

func init() {
    f := BDariya德拉伊耶.(*德拉伊耶DariyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dariya",
		TitleName: "德拉伊耶",
		TitleCode: "b_dariya",
	}
}
