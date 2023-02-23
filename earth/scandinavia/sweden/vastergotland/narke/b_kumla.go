package narke

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库姆拉KumlaBarony struct {
	feud.BaseBarony
}

var BKumla库姆拉 feud.Barony = &库姆拉KumlaBarony{}

func init() {
	f := BKumla库姆拉.(*库姆拉KumlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumla",
		TitleName: "库姆拉",
		TitleCode: "b_kumla",
	}
}
