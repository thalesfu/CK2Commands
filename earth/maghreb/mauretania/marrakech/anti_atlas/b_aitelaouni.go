package anti_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伊特阿乌里AitelaouniBarony struct {
	feud.BaseBarony
}

var BAitelaouni阿伊特阿乌里 feud.Barony = &阿伊特阿乌里AitelaouniBarony{}

func init() {
    f := BAitelaouni阿伊特阿乌里.(*阿伊特阿乌里AitelaouniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aitelaouni",
		TitleName: "阿伊特阿乌里",
		TitleCode: "b_aitelaouni",
	}
}
