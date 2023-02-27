package dorostotum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁希RusiBarony struct {
	feud.BaseBarony
}

var BRusi鲁希 feud.Barony = &鲁希RusiBarony{}

func init() {
    f := BRusi鲁希.(*鲁希RusiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rusi",
		TitleName: "鲁希",
		TitleCode: "b_rusi",
	}
}
