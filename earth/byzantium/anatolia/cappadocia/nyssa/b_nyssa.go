package nyssa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼撒NyssaBarony struct {
	feud.BaseBarony
}

var BNyssa尼撒 feud.Barony = &尼撒NyssaBarony{}

func init() {
    f := BNyssa尼撒.(*尼撒NyssaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyssa",
		TitleName: "尼撒",
		TitleCode: "b_nyssa",
	}
}
