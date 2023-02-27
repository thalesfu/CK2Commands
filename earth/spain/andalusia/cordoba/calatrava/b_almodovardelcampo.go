package calatrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔莫多瓦尔德尔坎波AlmodovardelcampoBarony struct {
	feud.BaseBarony
}

var BAlmodovardelcampo阿尔莫多瓦尔德尔坎波 feud.Barony = &阿尔莫多瓦尔德尔坎波AlmodovardelcampoBarony{}

func init() {
    f := BAlmodovardelcampo阿尔莫多瓦尔德尔坎波.(*阿尔莫多瓦尔德尔坎波AlmodovardelcampoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almodovardelcampo",
		TitleName: "阿尔莫多瓦尔德尔坎波",
		TitleCode: "b_almodovardelcampo",
	}
}
