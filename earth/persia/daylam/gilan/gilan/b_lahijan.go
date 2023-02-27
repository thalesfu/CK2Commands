package gilan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉希詹LahijanBarony struct {
	feud.BaseBarony
}

var BLahijan拉希詹 feud.Barony = &拉希詹LahijanBarony{}

func init() {
    f := BLahijan拉希詹.(*拉希詹LahijanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lahijan",
		TitleName: "拉希詹",
		TitleCode: "b_lahijan",
	}
}
