package gent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科特赖克KortrijkBarony struct {
	feud.BaseBarony
}

var BKortrijk科特赖克 feud.Barony = &科特赖克KortrijkBarony{}

func init() {
	f := BKortrijk科特赖克.(*科特赖克KortrijkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kortrijk",
		TitleName: "科特赖克",
		TitleCode: "b_kortrijk",
	}
}
