package dunhuang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 月牙泉YueyaquanBarony struct {
	feud.BaseBarony
}

var BYueyaquan月牙泉 feud.Barony = &月牙泉YueyaquanBarony{}

func init() {
    f := BYueyaquan月牙泉.(*月牙泉YueyaquanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yueyaquan",
		TitleName: "月牙泉",
		TitleCode: "b_yueyaquan",
	}
}
