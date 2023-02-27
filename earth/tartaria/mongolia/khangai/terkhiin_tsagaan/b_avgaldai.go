package terkhiin_tsagaan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布噶勒岱AvgaldaiBarony struct {
	feud.BaseBarony
}

var BAvgaldai阿布噶勒岱 feud.Barony = &阿布噶勒岱AvgaldaiBarony{}

func init() {
    f := BAvgaldai阿布噶勒岱.(*阿布噶勒岱AvgaldaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "avgaldai",
		TitleName: "阿布噶勒岱",
		TitleCode: "b_avgaldai",
	}
}
