package penthievre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙孔图尔MonkontourBarony struct {
	feud.BaseBarony
}

var BMonkontour蒙孔图尔 feud.Barony = &蒙孔图尔MonkontourBarony{}

func init() {
    f := BMonkontour蒙孔图尔.(*蒙孔图尔MonkontourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monkontour",
		TitleName: "蒙孔图尔",
		TitleCode: "b_monkontour",
	}
}
