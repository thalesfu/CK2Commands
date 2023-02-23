package taradavadi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般荼罗补罗PandharpurBarony struct {
	feud.BaseBarony
}

var BPandharpur般荼罗补罗 feud.Barony = &般荼罗补罗PandharpurBarony{}

func init() {
	f := BPandharpur般荼罗补罗.(*般荼罗补罗PandharpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pandharpur",
		TitleName: "般荼罗补罗",
		TitleCode: "b_pandharpur",
	}
}
