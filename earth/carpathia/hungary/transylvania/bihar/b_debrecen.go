package bihar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德布勒森DebrecenBarony struct {
	feud.BaseBarony
}

var BDebrecen德布勒森 feud.Barony = &德布勒森DebrecenBarony{}

func init() {
    f := BDebrecen德布勒森.(*德布勒森DebrecenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "debrecen",
		TitleName: "德布勒森",
		TitleCode: "b_debrecen",
	}
}
