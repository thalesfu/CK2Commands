package gafsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈马HammaBarony struct {
	feud.BaseBarony
}

var BHamma哈马 feud.Barony = &哈马HammaBarony{}

func init() {
    f := BHamma哈马.(*哈马HammaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamma",
		TitleName: "哈马",
		TitleCode: "b_hamma",
	}
}
