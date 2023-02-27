package dodugu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴马科BamakoBarony struct {
	feud.BaseBarony
}

var BBamako巴马科 feud.Barony = &巴马科BamakoBarony{}

func init() {
    f := BBamako巴马科.(*巴马科BamakoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bamako",
		TitleName: "巴马科",
		TitleCode: "b_bamako",
	}
}
