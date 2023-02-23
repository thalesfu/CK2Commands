package bannu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加拉伯克KalabaghBarony struct {
	feud.BaseBarony
}

var BKalabagh加拉伯克 feud.Barony = &加拉伯克KalabaghBarony{}

func init() {
	f := BKalabagh加拉伯克.(*加拉伯克KalabaghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalabagh",
		TitleName: "加拉伯克",
		TitleCode: "b_kalabagh",
	}
}
