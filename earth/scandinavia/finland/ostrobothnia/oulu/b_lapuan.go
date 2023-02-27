package oulu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉普阿LapuanBarony struct {
	feud.BaseBarony
}

var BLapuan拉普阿 feud.Barony = &拉普阿LapuanBarony{}

func init() {
    f := BLapuan拉普阿.(*拉普阿LapuanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lapuan",
		TitleName: "拉普阿",
		TitleCode: "b_lapuan",
	}
}
