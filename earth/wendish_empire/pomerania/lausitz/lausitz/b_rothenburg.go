package lausitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗滕堡RothenburgBarony struct {
	feud.BaseBarony
}

var BRothenburg罗滕堡 feud.Barony = &罗滕堡RothenburgBarony{}

func init() {
    f := BRothenburg罗滕堡.(*罗滕堡RothenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rothenburg",
		TitleName: "罗滕堡",
		TitleCode: "b_rothenburg",
	}
}
