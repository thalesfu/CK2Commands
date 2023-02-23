package meath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙斯KnowthBarony struct {
	feud.BaseBarony
}

var BKnowth瑙斯 feud.Barony = &瑙斯KnowthBarony{}

func init() {
	f := BKnowth瑙斯.(*瑙斯KnowthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "knowth",
		TitleName: "瑙斯",
		TitleCode: "b_knowth",
	}
}
