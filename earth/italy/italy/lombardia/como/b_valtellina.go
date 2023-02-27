package como

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔泰利纳ValtellinaBarony struct {
	feud.BaseBarony
}

var BValtellina瓦尔泰利纳 feud.Barony = &瓦尔泰利纳ValtellinaBarony{}

func init() {
    f := BValtellina瓦尔泰利纳.(*瓦尔泰利纳ValtellinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valtellina",
		TitleName: "瓦尔泰利纳",
		TitleCode: "b_valtellina",
	}
}
