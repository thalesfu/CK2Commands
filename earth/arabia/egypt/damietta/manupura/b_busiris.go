package manupura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布西里斯BusirisBarony struct {
	feud.BaseBarony
}

var BBusiris布西里斯 feud.Barony = &布西里斯BusirisBarony{}

func init() {
	f := BBusiris布西里斯.(*布西里斯BusirisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "busiris",
		TitleName: "布西里斯",
		TitleCode: "b_busiris",
	}
}
