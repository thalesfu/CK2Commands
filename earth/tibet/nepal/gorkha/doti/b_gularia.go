package doti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古拉里亚GulariaBarony struct {
	feud.BaseBarony
}

var BGularia古拉里亚 feud.Barony = &古拉里亚GulariaBarony{}

func init() {
	f := BGularia古拉里亚.(*古拉里亚GulariaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gularia",
		TitleName: "古拉里亚",
		TitleCode: "b_gularia",
	}
}
