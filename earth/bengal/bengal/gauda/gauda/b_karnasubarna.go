package gauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 羯罗拏苏伐剌那KarnasubarnaBarony struct {
	feud.BaseBarony
}

var BKarnasubarna羯罗拏苏伐剌那 feud.Barony = &羯罗拏苏伐剌那KarnasubarnaBarony{}

func init() {
	f := BKarnasubarna羯罗拏苏伐剌那.(*羯罗拏苏伐剌那KarnasubarnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karnasubarna",
		TitleName: "羯罗拏苏伐剌那",
		TitleCode: "b_karnasubarna",
	}
}
