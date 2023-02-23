package herat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库什克KushkBarony struct {
	feud.BaseBarony
}

var BKushk库什克 feud.Barony = &库什克KushkBarony{}

func init() {
	f := BKushk库什克.(*库什克KushkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kushk",
		TitleName: "库什克",
		TitleCode: "b_kushk",
	}
}
