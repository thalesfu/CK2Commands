package kremenchuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛玛迪奇MamadichBarony struct {
	feud.BaseBarony
}

var BMamadich玛玛迪奇 feud.Barony = &玛玛迪奇MamadichBarony{}

func init() {
    f := BMamadich玛玛迪奇.(*玛玛迪奇MamadichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mamadich",
		TitleName: "玛玛迪奇",
		TitleCode: "b_mamadich",
	}
}
