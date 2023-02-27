package lisboa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托马尔TomarBarony struct {
	feud.BaseBarony
}

var BTomar托马尔 feud.Barony = &托马尔TomarBarony{}

func init() {
    f := BTomar托马尔.(*托马尔TomarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tomar",
		TitleName: "托马尔",
		TitleCode: "b_tomar",
	}
}
