package strathearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克里夫CrieffBarony struct {
	feud.BaseBarony
}

var BCrieff克里夫 feud.Barony = &克里夫CrieffBarony{}

func init() {
    f := BCrieff克里夫.(*克里夫CrieffBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crieff",
		TitleName: "克里夫",
		TitleCode: "b_crieff",
	}
}
