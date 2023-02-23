package bourbon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙庞谢MontpensierBarony struct {
	feud.BaseBarony
}

var BMontpensier蒙庞谢 feud.Barony = &蒙庞谢MontpensierBarony{}

func init() {
	f := BMontpensier蒙庞谢.(*蒙庞谢MontpensierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montpensier",
		TitleName: "蒙庞谢",
		TitleCode: "b_montpensier",
	}
}
