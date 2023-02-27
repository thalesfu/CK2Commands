package madhupur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩度补罗MadhupurBarony struct {
	feud.BaseBarony
}

var BMadhupur摩度补罗 feud.Barony = &摩度补罗MadhupurBarony{}

func init() {
    f := BMadhupur摩度补罗.(*摩度补罗MadhupurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "madhupur",
		TitleName: "摩度补罗",
		TitleCode: "b_madhupur",
	}
}
