package galloway

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓弗里斯DumfriesBarony struct {
	feud.BaseBarony
}

var BDumfries邓弗里斯 feud.Barony = &邓弗里斯DumfriesBarony{}

func init() {
    f := BDumfries邓弗里斯.(*邓弗里斯DumfriesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dumfries",
		TitleName: "邓弗里斯",
		TitleCode: "b_dumfries",
	}
}
