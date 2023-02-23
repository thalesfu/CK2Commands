package mustang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珞LoBarony struct {
	feud.BaseBarony
}

var BLo珞 feud.Barony = &珞LoBarony{}

func init() {
	f := BLo珞.(*珞LoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lo",
		TitleName: "珞",
		TitleCode: "b_lo",
	}
}
