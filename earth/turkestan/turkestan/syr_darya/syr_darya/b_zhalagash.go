package syr_darya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎拉加什ZhalagashBarony struct {
	feud.BaseBarony
}

var BZhalagash扎拉加什 feud.Barony = &扎拉加什ZhalagashBarony{}

func init() {
    f := BZhalagash扎拉加什.(*扎拉加什ZhalagashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhalagash",
		TitleName: "扎拉加什",
		TitleCode: "b_zhalagash",
	}
}
