package bergenshus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卑尔根胡斯BergenhusBarony struct {
	feud.BaseBarony
}

var BBergenhus卑尔根胡斯 feud.Barony = &卑尔根胡斯BergenhusBarony{}

func init() {
    f := BBergenhus卑尔根胡斯.(*卑尔根胡斯BergenhusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bergenhus",
		TitleName: "卑尔根胡斯",
		TitleCode: "b_bergenhus",
	}
}
