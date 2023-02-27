package hail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈伊勒HailBarony struct {
	feud.BaseBarony
}

var BHail哈伊勒 feud.Barony = &哈伊勒HailBarony{}

func init() {
    f := BHail哈伊勒.(*哈伊勒HailBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hail",
		TitleName: "哈伊勒",
		TitleCode: "b_hail",
	}
}
