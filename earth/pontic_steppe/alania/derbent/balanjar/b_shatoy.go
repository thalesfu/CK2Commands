package balanjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙托伊ShatoyBarony struct {
	feud.BaseBarony
}

var BShatoy沙托伊 feud.Barony = &沙托伊ShatoyBarony{}

func init() {
    f := BShatoy沙托伊.(*沙托伊ShatoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shatoy",
		TitleName: "沙托伊",
		TitleCode: "b_shatoy",
	}
}
