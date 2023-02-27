package bayda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶什布姆YashbumBarony struct {
	feud.BaseBarony
}

var BYashbum耶什布姆 feud.Barony = &耶什布姆YashbumBarony{}

func init() {
    f := BYashbum耶什布姆.(*耶什布姆YashbumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yashbum",
		TitleName: "耶什布姆",
		TitleCode: "b_yashbum",
	}
}
