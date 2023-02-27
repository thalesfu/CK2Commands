package pamir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达拉乌特_库尔干Daroot_korgonBarony struct {
	feud.BaseBarony
}

var BDaroot_korgon达拉乌特_库尔干 feud.Barony = &达拉乌特_库尔干Daroot_korgonBarony{}

func init() {
    f := BDaroot_korgon达拉乌特_库尔干.(*达拉乌特_库尔干Daroot_korgonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daroot_korgon",
		TitleName: "达拉乌特_库尔干",
		TitleCode: "b_daroot_korgon",
	}
}
