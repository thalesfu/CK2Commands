package thuringen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗伊斯ReussBarony struct {
	feud.BaseBarony
}

var BReuss罗伊斯 feud.Barony = &罗伊斯ReussBarony{}

func init() {
	f := BReuss罗伊斯.(*罗伊斯ReussBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reuss",
		TitleName: "罗伊斯",
		TitleCode: "b_reuss",
	}
}
