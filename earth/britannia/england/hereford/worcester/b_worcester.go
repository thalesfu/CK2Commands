package worcester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伍斯特WorcesterBarony struct {
	feud.BaseBarony
}

var BWorcester伍斯特 feud.Barony = &伍斯特WorcesterBarony{}

func init() {
    f := BWorcester伍斯特.(*伍斯特WorcesterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "worcester",
		TitleName: "伍斯特",
		TitleCode: "b_worcester",
	}
}
