package pskov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈里斯托GoristoBarony struct {
	feud.BaseBarony
}

var BGoristo戈里斯托 feud.Barony = &戈里斯托GoristoBarony{}

func init() {
	f := BGoristo戈里斯托.(*戈里斯托GoristoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goristo",
		TitleName: "戈里斯托",
		TitleCode: "b_goristo",
	}
}
