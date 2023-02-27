package satakunta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波里PoriBarony struct {
	feud.BaseBarony
}

var BPori波里 feud.Barony = &波里PoriBarony{}

func init() {
    f := BPori波里.(*波里PoriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pori",
		TitleName: "波里",
		TitleCode: "b_pori",
	}
}
