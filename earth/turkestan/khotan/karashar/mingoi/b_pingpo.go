package mingoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 平坡PingpoBarony struct {
	feud.BaseBarony
}

var BPingpo平坡 feud.Barony = &平坡PingpoBarony{}

func init() {
    f := BPingpo平坡.(*平坡PingpoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pingpo",
		TitleName: "平坡",
		TitleCode: "b_pingpo",
	}
}
