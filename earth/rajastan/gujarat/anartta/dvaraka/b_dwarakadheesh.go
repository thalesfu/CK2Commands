package dvaraka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 堕罗迦地舍DwarakadheeshBarony struct {
	feud.BaseBarony
}

var BDwarakadheesh堕罗迦地舍 feud.Barony = &堕罗迦地舍DwarakadheeshBarony{}

func init() {
    f := BDwarakadheesh堕罗迦地舍.(*堕罗迦地舍DwarakadheeshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dwarakadheesh",
		TitleName: "堕罗迦地舍",
		TitleCode: "b_dwarakadheesh",
	}
}
