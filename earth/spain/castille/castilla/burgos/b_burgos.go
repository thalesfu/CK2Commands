package burgos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔戈斯BurgosBarony struct {
	feud.BaseBarony
}

var BBurgos布尔戈斯 feud.Barony = &布尔戈斯BurgosBarony{}

func init() {
    f := BBurgos布尔戈斯.(*布尔戈斯BurgosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burgos",
		TitleName: "布尔戈斯",
		TitleCode: "b_burgos",
	}
}
