package braganza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅略尔堡CastelomelhorBarony struct {
	feud.BaseBarony
}

var BCastelomelhor梅略尔堡 feud.Barony = &梅略尔堡CastelomelhorBarony{}

func init() {
    f := BCastelomelhor梅略尔堡.(*梅略尔堡CastelomelhorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castelomelhor",
		TitleName: "梅略尔堡",
		TitleCode: "b_castelomelhor",
	}
}
