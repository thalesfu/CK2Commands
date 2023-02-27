package hum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩奇PecBarony struct {
	feud.BaseBarony
}

var BPec佩奇 feud.Barony = &佩奇PecBarony{}

func init() {
    f := BPec佩奇.(*佩奇PecBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pec",
		TitleName: "佩奇",
		TitleCode: "b_pec",
	}
}
