package vendome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙图瓦尔MontoireBarony struct {
	feud.BaseBarony
}

var BMontoire蒙图瓦尔 feud.Barony = &蒙图瓦尔MontoireBarony{}

func init() {
	f := BMontoire蒙图瓦尔.(*蒙图瓦尔MontoireBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montoire",
		TitleName: "蒙图瓦尔",
		TitleCode: "b_montoire",
	}
}
