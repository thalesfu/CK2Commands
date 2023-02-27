package khinjali_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鸠尼案KunianBarony struct {
	feud.BaseBarony
}

var BKunian鸠尼案 feud.Barony = &鸠尼案KunianBarony{}

func init() {
    f := BKunian鸠尼案.(*鸠尼案KunianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunian",
		TitleName: "鸠尼案",
		TitleCode: "b_kunian",
	}
}
