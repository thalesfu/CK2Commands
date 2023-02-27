package vasyugan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦休甘VasyuganBarony struct {
	feud.BaseBarony
}

var BVasyugan瓦休甘 feud.Barony = &瓦休甘VasyuganBarony{}

func init() {
    f := BVasyugan瓦休甘.(*瓦休甘VasyuganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vasyugan",
		TitleName: "瓦休甘",
		TitleCode: "b_vasyugan",
	}
}
