package sarasvati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多罗因TarainBarony struct {
	feud.BaseBarony
}

var BTarain多罗因 feud.Barony = &多罗因TarainBarony{}

func init() {
	f := BTarain多罗因.(*多罗因TarainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarain",
		TitleName: "多罗因",
		TitleCode: "b_tarain",
	}
}
