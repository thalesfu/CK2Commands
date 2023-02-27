package braunschweig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑尔姆施泰特HelmstedtBarony struct {
	feud.BaseBarony
}

var BHelmstedt黑尔姆施泰特 feud.Barony = &黑尔姆施泰特HelmstedtBarony{}

func init() {
    f := BHelmstedt黑尔姆施泰特.(*黑尔姆施泰特HelmstedtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "helmstedt",
		TitleName: "黑尔姆施泰特",
		TitleCode: "b_helmstedt",
	}
}
