package lakhnau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑帝力迦SatrikhBarony struct {
	feud.BaseBarony
}

var BSatrikh娑帝力迦 feud.Barony = &娑帝力迦SatrikhBarony{}

func init() {
    f := BSatrikh娑帝力迦.(*娑帝力迦SatrikhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satrikh",
		TitleName: "娑帝力迦",
		TitleCode: "b_satrikh",
	}
}
