package medog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旁辛PhomshenBarony struct {
	feud.BaseBarony
}

var BPhomshen旁辛 feud.Barony = &旁辛PhomshenBarony{}

func init() {
    f := BPhomshen旁辛.(*旁辛PhomshenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phomshen",
		TitleName: "旁辛",
		TitleCode: "b_phomshen",
	}
}
