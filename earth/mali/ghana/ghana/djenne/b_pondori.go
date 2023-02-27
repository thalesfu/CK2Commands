package djenne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旁多里PondoriBarony struct {
	feud.BaseBarony
}

var BPondori旁多里 feud.Barony = &旁多里PondoriBarony{}

func init() {
    f := BPondori旁多里.(*旁多里PondoriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pondori",
		TitleName: "旁多里",
		TitleCode: "b_pondori",
	}
}
