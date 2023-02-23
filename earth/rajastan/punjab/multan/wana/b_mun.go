package wana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙河MunBarony struct {
	feud.BaseBarony
}

var BMun蒙河 feud.Barony = &蒙河MunBarony{}

func init() {
	f := BMun蒙河.(*蒙河MunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mun",
		TitleName: "蒙河",
		TitleCode: "b_mun",
	}
}
