package telemark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尔尤尔SeljordBarony struct {
	feud.BaseBarony
}

var BSeljord塞尔尤尔 feud.Barony = &塞尔尤尔SeljordBarony{}

func init() {
    f := BSeljord塞尔尤尔.(*塞尔尤尔SeljordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seljord",
		TitleName: "塞尔尤尔",
		TitleCode: "b_seljord",
	}
}
