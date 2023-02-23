package cairo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙尔马迪贝尼所罗蒙MerimdabenisalamaBarony struct {
	feud.BaseBarony
}

var BMerimdabenisalama蒙尔马迪贝尼所罗蒙 feud.Barony = &蒙尔马迪贝尼所罗蒙MerimdabenisalamaBarony{}

func init() {
	f := BMerimdabenisalama蒙尔马迪贝尼所罗蒙.(*蒙尔马迪贝尼所罗蒙MerimdabenisalamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merimdabenisalama",
		TitleName: "蒙尔马迪贝尼所罗蒙",
		TitleCode: "b_merimdabenisalama",
	}
}
