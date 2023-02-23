package galaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 万恰OanceaBarony struct {
	feud.BaseBarony
}

var BOancea万恰 feud.Barony = &万恰OanceaBarony{}

func init() {
	f := BOancea万恰.(*万恰OanceaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oancea",
		TitleName: "万恰",
		TitleCode: "b_oancea",
	}
}
