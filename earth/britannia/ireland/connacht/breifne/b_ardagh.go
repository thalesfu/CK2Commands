package breifne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿达ArdaghBarony struct {
	feud.BaseBarony
}

var BArdagh阿达 feud.Barony = &阿达ArdaghBarony{}

func init() {
    f := BArdagh阿达.(*阿达ArdaghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ardagh",
		TitleName: "阿达",
		TitleCode: "b_ardagh",
	}
}
