package juyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏泊淖尔SubozhuoerBarony struct {
	feud.BaseBarony
}

var BSubozhuoer苏泊淖尔 feud.Barony = &苏泊淖尔SubozhuoerBarony{}

func init() {
    f := BSubozhuoer苏泊淖尔.(*苏泊淖尔SubozhuoerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "subozhuoer",
		TitleName: "苏泊淖尔",
		TitleCode: "b_subozhuoer",
	}
}
