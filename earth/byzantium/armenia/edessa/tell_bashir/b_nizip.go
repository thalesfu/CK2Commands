package tell_bashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼济普NizipBarony struct {
	feud.BaseBarony
}

var BNizip尼济普 feud.Barony = &尼济普NizipBarony{}

func init() {
    f := BNizip尼济普.(*尼济普NizipBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nizip",
		TitleName: "尼济普",
		TitleCode: "b_nizip",
	}
}
