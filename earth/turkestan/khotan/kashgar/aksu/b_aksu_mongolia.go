package aksu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克苏Aksu_mongoliaBarony struct {
	feud.BaseBarony
}

var BAksu_mongolia阿克苏 feud.Barony = &阿克苏Aksu_mongoliaBarony{}

func init() {
    f := BAksu_mongolia阿克苏.(*阿克苏Aksu_mongoliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aksu_mongolia",
		TitleName: "阿克苏",
		TitleCode: "b_aksu_mongolia",
	}
}
