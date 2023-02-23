package almansa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔瓦塞特AlbaceteBarony struct {
	feud.BaseBarony
}

var BAlbacete阿尔瓦塞特 feud.Barony = &阿尔瓦塞特AlbaceteBarony{}

func init() {
	f := BAlbacete阿尔瓦塞特.(*阿尔瓦塞特AlbaceteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albacete",
		TitleName: "阿尔瓦塞特",
		TitleCode: "b_albacete",
	}
}
