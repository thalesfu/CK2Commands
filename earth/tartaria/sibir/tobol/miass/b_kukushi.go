package miass

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库库希KukushiBarony struct {
	feud.BaseBarony
}

var BKukushi库库希 feud.Barony = &库库希KukushiBarony{}

func init() {
	f := BKukushi库库希.(*库库希KukushiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kukushi",
		TitleName: "库库希",
		TitleCode: "b_kukushi",
	}
}
