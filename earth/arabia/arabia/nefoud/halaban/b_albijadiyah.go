package halaban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比贾迪亚AlbijadiyahBarony struct {
	feud.BaseBarony
}

var BAlbijadiyah比贾迪亚 feud.Barony = &比贾迪亚AlbijadiyahBarony{}

func init() {
    f := BAlbijadiyah比贾迪亚.(*比贾迪亚AlbijadiyahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albijadiyah",
		TitleName: "比贾迪亚",
		TitleCode: "b_albijadiyah",
	}
}
