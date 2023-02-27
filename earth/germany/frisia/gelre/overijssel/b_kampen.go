package overijssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎彭KampenBarony struct {
	feud.BaseBarony
}

var BKampen坎彭 feud.Barony = &坎彭KampenBarony{}

func init() {
    f := BKampen坎彭.(*坎彭KampenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kampen",
		TitleName: "坎彭",
		TitleCode: "b_kampen",
	}
}
