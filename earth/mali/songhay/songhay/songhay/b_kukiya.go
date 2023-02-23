package songhay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库基亚KukiyaBarony struct {
	feud.BaseBarony
}

var BKukiya库基亚 feud.Barony = &库基亚KukiyaBarony{}

func init() {
	f := BKukiya库基亚.(*库基亚KukiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kukiya",
		TitleName: "库基亚",
		TitleCode: "b_kukiya",
	}
}
