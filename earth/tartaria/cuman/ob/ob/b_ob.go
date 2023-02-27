package ob

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鄂毕ObBarony struct {
	feud.BaseBarony
}

var BOb鄂毕 feud.Barony = &鄂毕ObBarony{}

func init() {
    f := BOb鄂毕.(*鄂毕ObBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ob",
		TitleName: "鄂毕",
		TitleCode: "b_ob",
	}
}
