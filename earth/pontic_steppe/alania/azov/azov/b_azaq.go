package azov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿扎克AzaqBarony struct {
	feud.BaseBarony
}

var BAzaq阿扎克 feud.Barony = &阿扎克AzaqBarony{}

func init() {
    f := BAzaq阿扎克.(*阿扎克AzaqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azaq",
		TitleName: "阿扎克",
		TitleCode: "b_azaq",
	}
}
