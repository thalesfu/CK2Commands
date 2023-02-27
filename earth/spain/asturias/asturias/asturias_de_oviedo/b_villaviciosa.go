package asturias_de_oviedo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比利亚维西奥萨VillaviciosaBarony struct {
	feud.BaseBarony
}

var BVillaviciosa比利亚维西奥萨 feud.Barony = &比利亚维西奥萨VillaviciosaBarony{}

func init() {
    f := BVillaviciosa比利亚维西奥萨.(*比利亚维西奥萨VillaviciosaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villaviciosa",
		TitleName: "比利亚维西奥萨",
		TitleCode: "b_villaviciosa",
	}
}
