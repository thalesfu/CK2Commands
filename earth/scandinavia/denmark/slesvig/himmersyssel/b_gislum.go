package himmersyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉斯卢姆GislumBarony struct {
	feud.BaseBarony
}

var BGislum吉斯卢姆 feud.Barony = &吉斯卢姆GislumBarony{}

func init() {
    f := BGislum吉斯卢姆.(*吉斯卢姆GislumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gislum",
		TitleName: "吉斯卢姆",
		TitleCode: "b_gislum",
	}
}
