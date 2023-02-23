package lopnor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 七屯QitunBarony struct {
	feud.BaseBarony
}

var BQitun七屯 feud.Barony = &七屯QitunBarony{}

func init() {
	f := BQitun七屯.(*七屯QitunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qitun",
		TitleName: "七屯",
		TitleCode: "b_qitun",
	}
}
