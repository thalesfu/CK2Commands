package constantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科贾拉克CogealacBarony struct {
	feud.BaseBarony
}

var BCogealac科贾拉克 feud.Barony = &科贾拉克CogealacBarony{}

func init() {
	f := BCogealac科贾拉克.(*科贾拉克CogealacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cogealac",
		TitleName: "科贾拉克",
		TitleCode: "b_cogealac",
	}
}
