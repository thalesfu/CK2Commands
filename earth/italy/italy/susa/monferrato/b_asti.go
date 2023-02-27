package monferrato

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯蒂AstiBarony struct {
	feud.BaseBarony
}

var BAsti阿斯蒂 feud.Barony = &阿斯蒂AstiBarony{}

func init() {
    f := BAsti阿斯蒂.(*阿斯蒂AstiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asti",
		TitleName: "阿斯蒂",
		TitleCode: "b_asti",
	}
}
