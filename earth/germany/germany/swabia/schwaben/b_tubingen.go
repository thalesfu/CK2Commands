package schwaben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂宾根TubingenBarony struct {
	feud.BaseBarony
}

var BTubingen蒂宾根 feud.Barony = &蒂宾根TubingenBarony{}

func init() {
    f := BTubingen蒂宾根.(*蒂宾根TubingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tubingen",
		TitleName: "蒂宾根",
		TitleCode: "b_tubingen",
	}
}
