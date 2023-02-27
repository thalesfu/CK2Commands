package kiranapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 补耆DurgiBarony struct {
	feud.BaseBarony
}

var BDurgi补耆 feud.Barony = &补耆DurgiBarony{}

func init() {
    f := BDurgi补耆.(*补耆DurgiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durgi",
		TitleName: "补耆",
		TitleCode: "b_durgi",
	}
}
