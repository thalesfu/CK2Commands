package dithmarschen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦肯WackenBarony struct {
	feud.BaseBarony
}

var BWacken瓦肯 feud.Barony = &瓦肯WackenBarony{}

func init() {
    f := BWacken瓦肯.(*瓦肯WackenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wacken",
		TitleName: "瓦肯",
		TitleCode: "b_wacken",
	}
}
