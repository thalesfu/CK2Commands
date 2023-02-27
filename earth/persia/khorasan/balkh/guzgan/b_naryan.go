package guzgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那里杨NaryanBarony struct {
	feud.BaseBarony
}

var BNaryan那里杨 feud.Barony = &那里杨NaryanBarony{}

func init() {
    f := BNaryan那里杨.(*那里杨NaryanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naryan",
		TitleName: "那里杨",
		TitleCode: "b_naryan",
	}
}
