package nedong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孜塘ZetangBarony struct {
	feud.BaseBarony
}

var BZetang孜塘 feud.Barony = &孜塘ZetangBarony{}

func init() {
	f := BZetang孜塘.(*孜塘ZetangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zetang",
		TitleName: "孜塘",
		TitleCode: "b_zetang",
	}
}
