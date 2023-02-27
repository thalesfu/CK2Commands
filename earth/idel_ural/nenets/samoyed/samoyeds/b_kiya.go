package samoyeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基亚KiyaBarony struct {
	feud.BaseBarony
}

var BKiya基亚 feud.Barony = &基亚KiyaBarony{}

func init() {
    f := BKiya基亚.(*基亚KiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiya",
		TitleName: "基亚",
		TitleCode: "b_kiya",
	}
}
