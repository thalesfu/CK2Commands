package lusignan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼奥尔NiortBarony struct {
	feud.BaseBarony
}

var BNiort尼奥尔 feud.Barony = &尼奥尔NiortBarony{}

func init() {
    f := BNiort尼奥尔.(*尼奥尔NiortBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niort",
		TitleName: "尼奥尔",
		TitleCode: "b_niort",
	}
}
