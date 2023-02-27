package merya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约什卡尔_奥拉YoshkarolaBarony struct {
	feud.BaseBarony
}

var BYoshkarola约什卡尔_奥拉 feud.Barony = &约什卡尔_奥拉YoshkarolaBarony{}

func init() {
    f := BYoshkarola约什卡尔_奥拉.(*约什卡尔_奥拉YoshkarolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yoshkarola",
		TitleName: "约什卡尔_奥拉",
		TitleCode: "b_yoshkarola",
	}
}
