package sunnlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛格韦德利PingvellirBarony struct {
	feud.BaseBarony
}

var BPingvellir辛格韦德利 feud.Barony = &辛格韦德利PingvellirBarony{}

func init() {
    f := BPingvellir辛格韦德利.(*辛格韦德利PingvellirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pingvellir",
		TitleName: "辛格韦德利",
		TitleCode: "b_pingvellir",
	}
}
