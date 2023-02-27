package lothian

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 悲谷圣地Stow_of_wedaleBarony struct {
	feud.BaseBarony
}

var BStow_of_wedale悲谷圣地 feud.Barony = &悲谷圣地Stow_of_wedaleBarony{}

func init() {
    f := BStow_of_wedale悲谷圣地.(*悲谷圣地Stow_of_wedaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stow_of_wedale",
		TitleName: "悲谷圣地",
		TitleCode: "b_stow_of_wedale",
	}
}
