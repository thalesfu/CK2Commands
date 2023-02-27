package furstenberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈斯拉赫HaslachBarony struct {
	feud.BaseBarony
}

var BHaslach哈斯拉赫 feud.Barony = &哈斯拉赫HaslachBarony{}

func init() {
    f := BHaslach哈斯拉赫.(*哈斯拉赫HaslachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haslach",
		TitleName: "哈斯拉赫",
		TitleCode: "b_haslach",
	}
}
