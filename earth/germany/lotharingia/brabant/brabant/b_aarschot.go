package brabant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔斯霍特AarschotBarony struct {
	feud.BaseBarony
}

var BAarschot阿尔斯霍特 feud.Barony = &阿尔斯霍特AarschotBarony{}

func init() {
    f := BAarschot阿尔斯霍特.(*阿尔斯霍特AarschotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aarschot",
		TitleName: "阿尔斯霍特",
		TitleCode: "b_aarschot",
	}
}
