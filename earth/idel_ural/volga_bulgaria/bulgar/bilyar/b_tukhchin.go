package bilyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图克钦TukhchinBarony struct {
	feud.BaseBarony
}

var BTukhchin图克钦 feud.Barony = &图克钦TukhchinBarony{}

func init() {
    f := BTukhchin图克钦.(*图克钦TukhchinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tukhchin",
		TitleName: "图克钦",
		TitleCode: "b_tukhchin",
	}
}
