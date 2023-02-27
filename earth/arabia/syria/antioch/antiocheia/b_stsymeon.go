package antiocheia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣西梅昂StsymeonBarony struct {
	feud.BaseBarony
}

var BStsymeon圣西梅昂 feud.Barony = &圣西梅昂StsymeonBarony{}

func init() {
    f := BStsymeon圣西梅昂.(*圣西梅昂StsymeonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stsymeon",
		TitleName: "圣西梅昂",
		TitleCode: "b_stsymeon",
	}
}
