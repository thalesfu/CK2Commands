package daura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 津德尔ZinderBarony struct {
	feud.BaseBarony
}

var BZinder津德尔 feud.Barony = &津德尔ZinderBarony{}

func init() {
    f := BZinder津德尔.(*津德尔ZinderBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zinder",
		TitleName: "津德尔",
		TitleCode: "b_zinder",
	}
}
