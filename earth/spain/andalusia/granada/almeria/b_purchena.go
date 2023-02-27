package almeria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普尔切纳PurchenaBarony struct {
	feud.BaseBarony
}

var BPurchena普尔切纳 feud.Barony = &普尔切纳PurchenaBarony{}

func init() {
    f := BPurchena普尔切纳.(*普尔切纳PurchenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "purchena",
		TitleName: "普尔切纳",
		TitleCode: "b_purchena",
	}
}
