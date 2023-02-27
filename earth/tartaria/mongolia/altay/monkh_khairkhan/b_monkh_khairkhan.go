package monkh_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙赫海尔汗Monkh_khairkhanBarony struct {
	feud.BaseBarony
}

var BMonkh_khairkhan蒙赫海尔汗 feud.Barony = &蒙赫海尔汗Monkh_khairkhanBarony{}

func init() {
    f := BMonkh_khairkhan蒙赫海尔汗.(*蒙赫海尔汗Monkh_khairkhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monkh_khairkhan",
		TitleName: "蒙赫海尔汗",
		TitleCode: "b_monkh_khairkhan",
	}
}
