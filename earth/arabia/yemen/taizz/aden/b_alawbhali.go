package aden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥扎利AlawbhaliBarony struct {
	feud.BaseBarony
}

var BAlawbhali奥扎利 feud.Barony = &奥扎利AlawbhaliBarony{}

func init() {
	f := BAlawbhali奥扎利.(*奥扎利AlawbhaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alawbhali",
		TitleName: "奥扎利",
		TitleCode: "b_alawbhali",
	}
}
