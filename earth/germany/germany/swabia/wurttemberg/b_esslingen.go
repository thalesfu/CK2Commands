package wurttemberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯林根EsslingenBarony struct {
	feud.BaseBarony
}

var BEsslingen埃斯林根 feud.Barony = &埃斯林根EsslingenBarony{}

func init() {
    f := BEsslingen埃斯林根.(*埃斯林根EsslingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "esslingen",
		TitleName: "埃斯林根",
		TitleCode: "b_esslingen",
	}
}
