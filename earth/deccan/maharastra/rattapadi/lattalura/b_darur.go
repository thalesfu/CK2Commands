package lattalura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达鲁尔DarurBarony struct {
	feud.BaseBarony
}

var BDarur达鲁尔 feud.Barony = &达鲁尔DarurBarony{}

func init() {
	f := BDarur达鲁尔.(*达鲁尔DarurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darur",
		TitleName: "达鲁尔",
		TitleCode: "b_darur",
	}
}
