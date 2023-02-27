package ashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆达特MouddatBarony struct {
	feud.BaseBarony
}

var BMouddat穆达特 feud.Barony = &穆达特MouddatBarony{}

func init() {
    f := BMouddat穆达特.(*穆达特MouddatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mouddat",
		TitleName: "穆达特",
		TitleCode: "b_mouddat",
	}
}
