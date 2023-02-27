package sarangpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曷罗阇姞利呬RajgarhBarony struct {
	feud.BaseBarony
}

var BRajgarh曷罗阇姞利呬 feud.Barony = &曷罗阇姞利呬RajgarhBarony{}

func init() {
    f := BRajgarh曷罗阇姞利呬.(*曷罗阇姞利呬RajgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajgarh",
		TitleName: "曷罗阇姞利呬",
		TitleCode: "b_rajgarh",
	}
}
