package zaslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎斯拉夫ZaslavBarony struct {
	feud.BaseBarony
}

var BZaslav扎斯拉夫 feud.Barony = &扎斯拉夫ZaslavBarony{}

func init() {
    f := BZaslav扎斯拉夫.(*扎斯拉夫ZaslavBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaslav",
		TitleName: "扎斯拉夫",
		TitleCode: "b_zaslav",
	}
}
