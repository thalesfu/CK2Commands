package yazd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎尔奇ZarchBarony struct {
	feud.BaseBarony
}

var BZarch扎尔奇 feud.Barony = &扎尔奇ZarchBarony{}

func init() {
    f := BZarch扎尔奇.(*扎尔奇ZarchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zarch",
		TitleName: "扎尔奇",
		TitleCode: "b_zarch",
	}
}
