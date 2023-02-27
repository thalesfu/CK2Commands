package tavasts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海科宁HaikonenBarony struct {
	feud.BaseBarony
}

var BHaikonen海科宁 feud.Barony = &海科宁HaikonenBarony{}

func init() {
    f := BHaikonen海科宁.(*海科宁HaikonenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haikonen",
		TitleName: "海科宁",
		TitleCode: "b_haikonen",
	}
}
