package vardhamana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇佉梨JakhriBarony struct {
	feud.BaseBarony
}

var BJakhri阇佉梨 feud.Barony = &阇佉梨JakhriBarony{}

func init() {
    f := BJakhri阇佉梨.(*阇佉梨JakhriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jakhri",
		TitleName: "阇佉梨",
		TitleCode: "b_jakhri",
	}
}
