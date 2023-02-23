package lykia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米拉MyraBarony struct {
	feud.BaseBarony
}

var BMyra米拉 feud.Barony = &米拉MyraBarony{}

func init() {
	f := BMyra米拉.(*米拉MyraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "myra",
		TitleName: "米拉",
		TitleCode: "b_myra",
	}
}
