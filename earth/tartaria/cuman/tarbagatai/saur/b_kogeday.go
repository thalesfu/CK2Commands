package saur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库库岱KogedayBarony struct {
	feud.BaseBarony
}

var BKogeday库库岱 feud.Barony = &库库岱KogedayBarony{}

func init() {
	f := BKogeday库库岱.(*库库岱KogedayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kogeday",
		TitleName: "库库岱",
		TitleCode: "b_kogeday",
	}
}
