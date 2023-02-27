package nandagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫拉瓦吉尔MulavagilBarony struct {
	feud.BaseBarony
}

var BMulavagil莫拉瓦吉尔 feud.Barony = &莫拉瓦吉尔MulavagilBarony{}

func init() {
    f := BMulavagil莫拉瓦吉尔.(*莫拉瓦吉尔MulavagilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mulavagil",
		TitleName: "莫拉瓦吉尔",
		TitleCode: "b_mulavagil",
	}
}
