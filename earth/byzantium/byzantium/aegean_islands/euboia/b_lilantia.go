package euboia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利兰迪亚LilantiaBarony struct {
	feud.BaseBarony
}

var BLilantia利兰迪亚 feud.Barony = &利兰迪亚LilantiaBarony{}

func init() {
    f := BLilantia利兰迪亚.(*利兰迪亚LilantiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lilantia",
		TitleName: "利兰迪亚",
		TitleCode: "b_lilantia",
	}
}
