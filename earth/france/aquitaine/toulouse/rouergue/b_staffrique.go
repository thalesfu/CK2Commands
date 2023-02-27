package rouergue

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣阿夫里克StaffriqueBarony struct {
	feud.BaseBarony
}

var BStaffrique圣阿夫里克 feud.Barony = &圣阿夫里克StaffriqueBarony{}

func init() {
    f := BStaffrique圣阿夫里克.(*圣阿夫里克StaffriqueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "staffrique",
		TitleName: "圣阿夫里克",
		TitleCode: "b_staffrique",
	}
}
