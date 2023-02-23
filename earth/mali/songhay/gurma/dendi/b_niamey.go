package dendi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼亚美NiameyBarony struct {
	feud.BaseBarony
}

var BNiamey尼亚美 feud.Barony = &尼亚美NiameyBarony{}

func init() {
	f := BNiamey尼亚美.(*尼亚美NiameyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niamey",
		TitleName: "尼亚美",
		TitleCode: "b_niamey",
	}
}
