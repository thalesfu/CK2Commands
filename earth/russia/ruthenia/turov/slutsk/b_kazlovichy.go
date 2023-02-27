package slutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科兹洛维奇KazlovichyBarony struct {
	feud.BaseBarony
}

var BKazlovichy科兹洛维奇 feud.Barony = &科兹洛维奇KazlovichyBarony{}

func init() {
    f := BKazlovichy科兹洛维奇.(*科兹洛维奇KazlovichyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kazlovichy",
		TitleName: "科兹洛维奇",
		TitleCode: "b_kazlovichy",
	}
}
