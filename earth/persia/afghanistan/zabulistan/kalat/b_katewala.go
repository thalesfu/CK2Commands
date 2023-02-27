package kalat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽谛伐罗KatewalaBarony struct {
	feud.BaseBarony
}

var BKatewala伽谛伐罗 feud.Barony = &伽谛伐罗KatewalaBarony{}

func init() {
    f := BKatewala伽谛伐罗.(*伽谛伐罗KatewalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "katewala",
		TitleName: "伽谛伐罗",
		TitleCode: "b_katewala",
	}
}
