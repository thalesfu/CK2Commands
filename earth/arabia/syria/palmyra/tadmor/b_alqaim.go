package tadmor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加伊姆AlqaimBarony struct {
	feud.BaseBarony
}

var BAlqaim加伊姆 feud.Barony = &加伊姆AlqaimBarony{}

func init() {
    f := BAlqaim加伊姆.(*加伊姆AlqaimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alqaim",
		TitleName: "加伊姆",
		TitleCode: "b_alqaim",
	}
}
