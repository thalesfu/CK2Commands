package vermandois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣康坦StquentinBarony struct {
	feud.BaseBarony
}

var BStquentin圣康坦 feud.Barony = &圣康坦StquentinBarony{}

func init() {
	f := BStquentin圣康坦.(*圣康坦StquentinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stquentin",
		TitleName: "圣康坦",
		TitleCode: "b_stquentin",
	}
}
