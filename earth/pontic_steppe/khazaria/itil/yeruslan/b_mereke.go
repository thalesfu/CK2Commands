package yeruslan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅列克MerekeBarony struct {
	feud.BaseBarony
}

var BMereke梅列克 feud.Barony = &梅列克MerekeBarony{}

func init() {
    f := BMereke梅列克.(*梅列克MerekeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mereke",
		TitleName: "梅列克",
		TitleCode: "b_mereke",
	}
}
