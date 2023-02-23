package nordgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特拉斯堡StrassburgBarony struct {
	feud.BaseBarony
}

var BStrassburg斯特拉斯堡 feud.Barony = &斯特拉斯堡StrassburgBarony{}

func init() {
	f := BStrassburg斯特拉斯堡.(*斯特拉斯堡StrassburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strassburg",
		TitleName: "斯特拉斯堡",
		TitleCode: "b_strassburg",
	}
}
