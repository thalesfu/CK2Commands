package dasapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兴具罗阇姞利呬HinglajgarhBarony struct {
	feud.BaseBarony
}

var BHinglajgarh兴具罗阇姞利呬 feud.Barony = &兴具罗阇姞利呬HinglajgarhBarony{}

func init() {
    f := BHinglajgarh兴具罗阇姞利呬.(*兴具罗阇姞利呬HinglajgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hinglajgarh",
		TitleName: "兴具罗阇姞利呬",
		TitleCode: "b_hinglajgarh",
	}
}
