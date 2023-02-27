package konjikala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢拉赫斯SarahsBarony struct {
	feud.BaseBarony
}

var BSarahs谢拉赫斯 feud.Barony = &谢拉赫斯SarahsBarony{}

func init() {
    f := BSarahs谢拉赫斯.(*谢拉赫斯SarahsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarahs",
		TitleName: "谢拉赫斯",
		TitleCode: "b_sarahs",
	}
}
