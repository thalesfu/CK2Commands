package adana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩普绥提亚MopsuestiaBarony struct {
	feud.BaseBarony
}

var BMopsuestia摩普绥提亚 feud.Barony = &摩普绥提亚MopsuestiaBarony{}

func init() {
    f := BMopsuestia摩普绥提亚.(*摩普绥提亚MopsuestiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mopsuestia",
		TitleName: "摩普绥提亚",
		TitleCode: "b_mopsuestia",
	}
}
