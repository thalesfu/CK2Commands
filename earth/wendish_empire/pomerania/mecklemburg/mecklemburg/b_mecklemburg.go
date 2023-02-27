package mecklemburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅克伦堡MecklemburgBarony struct {
	feud.BaseBarony
}

var BMecklemburg梅克伦堡 feud.Barony = &梅克伦堡MecklemburgBarony{}

func init() {
    f := BMecklemburg梅克伦堡.(*梅克伦堡MecklemburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mecklemburg",
		TitleName: "梅克伦堡",
		TitleCode: "b_mecklemburg",
	}
}
