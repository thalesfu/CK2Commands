package nikomedeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼科米底亚NikomedeiaBarony struct {
	feud.BaseBarony
}

var BNikomedeia尼科米底亚 feud.Barony = &尼科米底亚NikomedeiaBarony{}

func init() {
	f := BNikomedeia尼科米底亚.(*尼科米底亚NikomedeiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nikomedeia",
		TitleName: "尼科米底亚",
		TitleCode: "b_nikomedeia",
	}
}
