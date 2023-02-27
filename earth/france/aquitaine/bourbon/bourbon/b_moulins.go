package bourbon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆兰MoulinsBarony struct {
	feud.BaseBarony
}

var BMoulins穆兰 feud.Barony = &穆兰MoulinsBarony{}

func init() {
    f := BMoulins穆兰.(*穆兰MoulinsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moulins",
		TitleName: "穆兰",
		TitleCode: "b_moulins",
	}
}
