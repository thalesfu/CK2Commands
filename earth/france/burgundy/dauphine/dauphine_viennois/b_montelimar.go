package dauphine_viennois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特利马尔MontelimarBarony struct {
	feud.BaseBarony
}

var BMontelimar蒙特利马尔 feud.Barony = &蒙特利马尔MontelimarBarony{}

func init() {
    f := BMontelimar蒙特利马尔.(*蒙特利马尔MontelimarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montelimar",
		TitleName: "蒙特利马尔",
		TitleCode: "b_montelimar",
	}
}
