package madaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨哈布SahabBarony struct {
	feud.BaseBarony
}

var BSahab萨哈布 feud.Barony = &萨哈布SahabBarony{}

func init() {
    f := BSahab萨哈布.(*萨哈布SahabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sahab",
		TitleName: "萨哈布",
		TitleCode: "b_sahab",
	}
}
