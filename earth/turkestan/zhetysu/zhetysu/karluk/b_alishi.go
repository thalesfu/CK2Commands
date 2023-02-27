package karluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿利施AlishiBarony struct {
	feud.BaseBarony
}

var BAlishi阿利施 feud.Barony = &阿利施AlishiBarony{}

func init() {
    f := BAlishi阿利施.(*阿利施AlishiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alishi",
		TitleName: "阿利施",
		TitleCode: "b_alishi",
	}
}
