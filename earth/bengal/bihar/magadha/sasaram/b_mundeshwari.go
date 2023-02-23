package sasaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙提湿伐罗MundeshwariBarony struct {
	feud.BaseBarony
}

var BMundeshwari蒙提湿伐罗 feud.Barony = &蒙提湿伐罗MundeshwariBarony{}

func init() {
	f := BMundeshwari蒙提湿伐罗.(*蒙提湿伐罗MundeshwariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mundeshwari",
		TitleName: "蒙提湿伐罗",
		TitleCode: "b_mundeshwari",
	}
}
