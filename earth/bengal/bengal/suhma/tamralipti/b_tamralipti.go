package tamralipti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多摩梨帝TamraliptiBarony struct {
	feud.BaseBarony
}

var BTamralipti多摩梨帝 feud.Barony = &多摩梨帝TamraliptiBarony{}

func init() {
	f := BTamralipti多摩梨帝.(*多摩梨帝TamraliptiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamralipti",
		TitleName: "多摩梨帝",
		TitleCode: "b_tamralipti",
	}
}
