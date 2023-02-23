package acre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克AcreBarony struct {
	feud.BaseBarony
}

var BAcre阿克 feud.Barony = &阿克AcreBarony{}

func init() {
	f := BAcre阿克.(*阿克AcreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "acre",
		TitleName: "阿克",
		TitleCode: "b_acre",
	}
}
