package uwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆哈拉格MuharraqBarony struct {
	feud.BaseBarony
}

var BMuharraq穆哈拉格 feud.Barony = &穆哈拉格MuharraqBarony{}

func init() {
    f := BMuharraq穆哈拉格.(*穆哈拉格MuharraqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muharraq",
		TitleName: "穆哈拉格",
		TitleCode: "b_muharraq",
	}
}
