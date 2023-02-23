package urgell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔赫尔UrgellBarony struct {
	feud.BaseBarony
}

var BUrgell乌尔赫尔 feud.Barony = &乌尔赫尔UrgellBarony{}

func init() {
	f := BUrgell乌尔赫尔.(*乌尔赫尔UrgellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urgell",
		TitleName: "乌尔赫尔",
		TitleCode: "b_urgell",
	}
}
