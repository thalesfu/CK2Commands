package kalat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆明补罗BhaminpurBarony struct {
	feud.BaseBarony
}

var BBhaminpur婆明补罗 feud.Barony = &婆明补罗BhaminpurBarony{}

func init() {
    f := BBhaminpur婆明补罗.(*婆明补罗BhaminpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhaminpur",
		TitleName: "婆明补罗",
		TitleCode: "b_bhaminpur",
	}
}
