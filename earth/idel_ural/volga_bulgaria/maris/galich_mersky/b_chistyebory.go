package galich_mersky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇斯特耶博雷ChistyeboryBarony struct {
	feud.BaseBarony
}

var BChistyebory奇斯特耶博雷 feud.Barony = &奇斯特耶博雷ChistyeboryBarony{}

func init() {
    f := BChistyebory奇斯特耶博雷.(*奇斯特耶博雷ChistyeboryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chistyebory",
		TitleName: "奇斯特耶博雷",
		TitleCode: "b_chistyebory",
	}
}
