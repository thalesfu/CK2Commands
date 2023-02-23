package albarracin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙塔尔万MontalbanBarony struct {
	feud.BaseBarony
}

var BMontalban蒙塔尔万 feud.Barony = &蒙塔尔万MontalbanBarony{}

func init() {
	f := BMontalban蒙塔尔万.(*蒙塔尔万MontalbanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montalban",
		TitleName: "蒙塔尔万",
		TitleCode: "b_montalban",
	}
}
