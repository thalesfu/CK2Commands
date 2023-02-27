package kartli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 第比利斯TbilisiBarony struct {
	feud.BaseBarony
}

var BTbilisi第比利斯 feud.Barony = &第比利斯TbilisiBarony{}

func init() {
    f := BTbilisi第比利斯.(*第比利斯TbilisiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tbilisi",
		TitleName: "第比利斯",
		TitleCode: "b_tbilisi",
	}
}
