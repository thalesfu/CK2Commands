package izborsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡姆诺KamnoBarony struct {
	feud.BaseBarony
}

var BKamno卡姆诺 feud.Barony = &卡姆诺KamnoBarony{}

func init() {
	f := BKamno卡姆诺.(*卡姆诺KamnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamno",
		TitleName: "卡姆诺",
		TitleCode: "b_kamno",
	}
}
