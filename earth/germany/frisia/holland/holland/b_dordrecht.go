package holland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多德雷赫特DordrechtBarony struct {
	feud.BaseBarony
}

var BDordrecht多德雷赫特 feud.Barony = &多德雷赫特DordrechtBarony{}

func init() {
    f := BDordrecht多德雷赫特.(*多德雷赫特DordrechtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dordrecht",
		TitleName: "多德雷赫特",
		TitleCode: "b_dordrecht",
	}
}
