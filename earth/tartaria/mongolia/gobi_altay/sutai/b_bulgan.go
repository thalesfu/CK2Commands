package sutai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔干BulganBarony struct {
	feud.BaseBarony
}

var BBulgan布尔干 feud.Barony = &布尔干BulganBarony{}

func init() {
    f := BBulgan布尔干.(*布尔干BulganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bulgan",
		TitleName: "布尔干",
		TitleCode: "b_bulgan",
	}
}
