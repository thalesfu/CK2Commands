package alqusair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆鲁斯UmmrusBarony struct {
	feud.BaseBarony
}

var BUmmrus乌姆鲁斯 feud.Barony = &乌姆鲁斯UmmrusBarony{}

func init() {
	f := BUmmrus乌姆鲁斯.(*乌姆鲁斯UmmrusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ummrus",
		TitleName: "乌姆鲁斯",
		TitleCode: "b_ummrus",
	}
}
