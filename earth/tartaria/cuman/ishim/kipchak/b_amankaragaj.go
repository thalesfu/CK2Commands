package kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿曼卡拉盖AmankaragajBarony struct {
	feud.BaseBarony
}

var BAmankaragaj阿曼卡拉盖 feud.Barony = &阿曼卡拉盖AmankaragajBarony{}

func init() {
    f := BAmankaragaj阿曼卡拉盖.(*阿曼卡拉盖AmankaragajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amankaragaj",
		TitleName: "阿曼卡拉盖",
		TitleCode: "b_amankaragaj",
	}
}
