package reval

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕德斯PadesBarony struct {
	feud.BaseBarony
}

var BPades帕德斯 feud.Barony = &帕德斯PadesBarony{}

func init() {
	f := BPades帕德斯.(*帕德斯PadesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pades",
		TitleName: "帕德斯",
		TitleCode: "b_pades",
	}
}
