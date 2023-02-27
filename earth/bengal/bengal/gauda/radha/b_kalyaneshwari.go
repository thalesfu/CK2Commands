package radha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦梨耶泥湿伐罗KalyaneshwariBarony struct {
	feud.BaseBarony
}

var BKalyaneshwari迦梨耶泥湿伐罗 feud.Barony = &迦梨耶泥湿伐罗KalyaneshwariBarony{}

func init() {
    f := BKalyaneshwari迦梨耶泥湿伐罗.(*迦梨耶泥湿伐罗KalyaneshwariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalyaneshwari",
		TitleName: "迦梨耶泥湿伐罗",
		TitleCode: "b_kalyaneshwari",
	}
}
