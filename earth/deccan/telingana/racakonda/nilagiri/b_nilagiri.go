package nilagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼罗耆厘NilagiriBarony struct {
	feud.BaseBarony
}

var BNilagiri尼罗耆厘 feud.Barony = &尼罗耆厘NilagiriBarony{}

func init() {
    f := BNilagiri尼罗耆厘.(*尼罗耆厘NilagiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nilagiri",
		TitleName: "尼罗耆厘",
		TitleCode: "b_nilagiri",
	}
}
