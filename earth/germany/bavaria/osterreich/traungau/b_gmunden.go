package traungau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格蒙登GmundenBarony struct {
	feud.BaseBarony
}

var BGmunden格蒙登 feud.Barony = &格蒙登GmundenBarony{}

func init() {
	f := BGmunden格蒙登.(*格蒙登GmundenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gmunden",
		TitleName: "格蒙登",
		TitleCode: "b_gmunden",
	}
}
