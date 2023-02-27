package upper_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布热格BrzegBarony struct {
	feud.BaseBarony
}

var BBrzeg布热格 feud.Barony = &布热格BrzegBarony{}

func init() {
    f := BBrzeg布热格.(*布热格BrzegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brzeg",
		TitleName: "布热格",
		TitleCode: "b_brzeg",
	}
}
