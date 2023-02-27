package venaissin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 教宗新堡ChateauneufdupapeBarony struct {
	feud.BaseBarony
}

var BChateauneufdupape教宗新堡 feud.Barony = &教宗新堡ChateauneufdupapeBarony{}

func init() {
    f := BChateauneufdupape教宗新堡.(*教宗新堡ChateauneufdupapeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chateauneufdupape",
		TitleName: "教宗新堡",
		TitleCode: "b_chateauneufdupape",
	}
}
