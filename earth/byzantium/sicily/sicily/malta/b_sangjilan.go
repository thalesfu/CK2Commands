package malta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳斯卡里奥SangjilanBarony struct {
	feud.BaseBarony
}

var BSangjilan纳斯卡里奥 feud.Barony = &纳斯卡里奥SangjilanBarony{}

func init() {
	f := BSangjilan纳斯卡里奥.(*纳斯卡里奥SangjilanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangjilan",
		TitleName: "纳斯卡里奥",
		TitleCode: "b_sangjilan",
	}
}
