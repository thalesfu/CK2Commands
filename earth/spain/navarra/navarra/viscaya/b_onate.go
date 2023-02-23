package viscaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尼亚蒂OnateBarony struct {
	feud.BaseBarony
}

var BOnate奥尼亚蒂 feud.Barony = &奥尼亚蒂OnateBarony{}

func init() {
	f := BOnate奥尼亚蒂.(*奥尼亚蒂OnateBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "onate",
		TitleName: "奥尼亚蒂",
		TitleCode: "b_onate",
	}
}
