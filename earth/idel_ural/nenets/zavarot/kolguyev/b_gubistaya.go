package kolguyev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古比斯塔亚GubistayaBarony struct {
	feud.BaseBarony
}

var BGubistaya古比斯塔亚 feud.Barony = &古比斯塔亚GubistayaBarony{}

func init() {
    f := BGubistaya古比斯塔亚.(*古比斯塔亚GubistayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gubistaya",
		TitleName: "古比斯塔亚",
		TitleCode: "b_gubistaya",
	}
}
