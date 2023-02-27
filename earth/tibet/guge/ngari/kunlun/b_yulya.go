package kunlun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 月牙YulyaBarony struct {
	feud.BaseBarony
}

var BYulya月牙 feud.Barony = &月牙YulyaBarony{}

func init() {
    f := BYulya月牙.(*月牙YulyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yulya",
		TitleName: "月牙",
		TitleCode: "b_yulya",
	}
}
