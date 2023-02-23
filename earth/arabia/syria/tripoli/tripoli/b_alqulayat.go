package tripoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古莱阿特AlqulayatBarony struct {
	feud.BaseBarony
}

var BAlqulayat古莱阿特 feud.Barony = &古莱阿特AlqulayatBarony{}

func init() {
	f := BAlqulayat古莱阿特.(*古莱阿特AlqulayatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alqulayat",
		TitleName: "古莱阿特",
		TitleCode: "b_alqulayat",
	}
}
