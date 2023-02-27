package urbino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特费尔特罗MontefeltroBarony struct {
	feud.BaseBarony
}

var BMontefeltro蒙特费尔特罗 feud.Barony = &蒙特费尔特罗MontefeltroBarony{}

func init() {
    f := BMontefeltro蒙特费尔特罗.(*蒙特费尔特罗MontefeltroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montefeltro",
		TitleName: "蒙特费尔特罗",
		TitleCode: "b_montefeltro",
	}
}
