package monreal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙雷阿尔MonrealBarony struct {
	feud.BaseBarony
}

var BMonreal蒙雷阿尔 feud.Barony = &蒙雷阿尔MonrealBarony{}

func init() {
    f := BMonreal蒙雷阿尔.(*蒙雷阿尔MonrealBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monreal",
		TitleName: "蒙雷阿尔",
		TitleCode: "b_monreal",
	}
}
