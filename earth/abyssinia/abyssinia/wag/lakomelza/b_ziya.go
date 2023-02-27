package lakomelza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹雅ZiyaBarony struct {
	feud.BaseBarony
}

var BZiya兹雅 feud.Barony = &兹雅ZiyaBarony{}

func init() {
    f := BZiya兹雅.(*兹雅ZiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ziya",
		TitleName: "兹雅",
		TitleCode: "b_ziya",
	}
}
