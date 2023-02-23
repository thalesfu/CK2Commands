package brescia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩斯基耶拉PeschieraBarony struct {
	feud.BaseBarony
}

var BPeschiera佩斯基耶拉 feud.Barony = &佩斯基耶拉PeschieraBarony{}

func init() {
	f := BPeschiera佩斯基耶拉.(*佩斯基耶拉PeschieraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "peschiera",
		TitleName: "佩斯基耶拉",
		TitleCode: "b_peschiera",
	}
}
