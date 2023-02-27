package akhmim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝萨BesaBarony struct {
	feud.BaseBarony
}

var BBesa贝萨 feud.Barony = &贝萨BesaBarony{}

func init() {
    f := BBesa贝萨.(*贝萨BesaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "besa",
		TitleName: "贝萨",
		TitleCode: "b_besa",
	}
}
