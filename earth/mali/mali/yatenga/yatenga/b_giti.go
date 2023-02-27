package yatenga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉提GitiBarony struct {
	feud.BaseBarony
}

var BGiti吉提 feud.Barony = &吉提GitiBarony{}

func init() {
    f := BGiti吉提.(*吉提GitiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "giti",
		TitleName: "吉提",
		TitleCode: "b_giti",
	}
}
