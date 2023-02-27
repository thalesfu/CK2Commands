package vesoul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕尔LureBarony struct {
	feud.BaseBarony
}

var BLure吕尔 feud.Barony = &吕尔LureBarony{}

func init() {
    f := BLure吕尔.(*吕尔LureBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lure",
		TitleName: "吕尔",
		TitleCode: "b_lure",
	}
}
