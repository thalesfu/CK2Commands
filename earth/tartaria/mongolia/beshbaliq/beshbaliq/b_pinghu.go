package beshbaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 平湖PinghuBarony struct {
	feud.BaseBarony
}

var BPinghu平湖 feud.Barony = &平湖PinghuBarony{}

func init() {
    f := BPinghu平湖.(*平湖PinghuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pinghu",
		TitleName: "平湖",
		TitleCode: "b_pinghu",
	}
}
