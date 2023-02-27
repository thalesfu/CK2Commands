package dvaraka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 堕罗迦DvarakaBarony struct {
	feud.BaseBarony
}

var BDvaraka堕罗迦 feud.Barony = &堕罗迦DvarakaBarony{}

func init() {
    f := BDvaraka堕罗迦.(*堕罗迦DvarakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dvaraka",
		TitleName: "堕罗迦",
		TitleCode: "b_dvaraka",
	}
}
