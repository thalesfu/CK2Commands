package chandax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡蒂亚SitiaBarony struct {
	feud.BaseBarony
}

var BSitia锡蒂亚 feud.Barony = &锡蒂亚SitiaBarony{}

func init() {
	f := BSitia锡蒂亚.(*锡蒂亚SitiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sitia",
		TitleName: "锡蒂亚",
		TitleCode: "b_sitia",
	}
}
