package navasarika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那婆萨利迦NavasarikaBarony struct {
	feud.BaseBarony
}

var BNavasarika那婆萨利迦 feud.Barony = &那婆萨利迦NavasarikaBarony{}

func init() {
	f := BNavasarika那婆萨利迦.(*那婆萨利迦NavasarikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "navasarika",
		TitleName: "那婆萨利迦",
		TitleCode: "b_navasarika",
	}
}
