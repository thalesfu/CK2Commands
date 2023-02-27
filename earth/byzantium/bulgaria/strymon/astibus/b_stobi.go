package astibus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯托比StobiBarony struct {
	feud.BaseBarony
}

var BStobi斯托比 feud.Barony = &斯托比StobiBarony{}

func init() {
    f := BStobi斯托比.(*斯托比StobiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stobi",
		TitleName: "斯托比",
		TitleCode: "b_stobi",
	}
}
