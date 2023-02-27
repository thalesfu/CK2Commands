package dvaraka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般跋BhanvadBarony struct {
	feud.BaseBarony
}

var BBhanvad般跋 feud.Barony = &般跋BhanvadBarony{}

func init() {
    f := BBhanvad般跋.(*般跋BhanvadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhanvad",
		TitleName: "般跋",
		TitleCode: "b_bhanvad",
	}
}
