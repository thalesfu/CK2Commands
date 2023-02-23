package angouleme

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昂古莱姆AngoulemeBarony struct {
	feud.BaseBarony
}

var BAngouleme昂古莱姆 feud.Barony = &昂古莱姆AngoulemeBarony{}

func init() {
	f := BAngouleme昂古莱姆.(*昂古莱姆AngoulemeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "angouleme",
		TitleName: "昂古莱姆",
		TitleCode: "b_angouleme",
	}
}
