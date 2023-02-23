package tadla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维泽古尔WazeqqurBarony struct {
	feud.BaseBarony
}

var BWazeqqur维泽古尔 feud.Barony = &维泽古尔WazeqqurBarony{}

func init() {
	f := BWazeqqur维泽古尔.(*维泽古尔WazeqqurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wazeqqur",
		TitleName: "维泽古尔",
		TitleCode: "b_wazeqqur",
	}
}
