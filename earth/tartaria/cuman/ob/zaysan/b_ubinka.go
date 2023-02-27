package zaysan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌宾卡UbinkaBarony struct {
	feud.BaseBarony
}

var BUbinka乌宾卡 feud.Barony = &乌宾卡UbinkaBarony{}

func init() {
    f := BUbinka乌宾卡.(*乌宾卡UbinkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ubinka",
		TitleName: "乌宾卡",
		TitleCode: "b_ubinka",
	}
}
