package wielunska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维隆WielunBarony struct {
	feud.BaseBarony
}

var BWielun维隆 feud.Barony = &维隆WielunBarony{}

func init() {
    f := BWielun维隆.(*维隆WielunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wielun",
		TitleName: "维隆",
		TitleCode: "b_wielun",
	}
}
