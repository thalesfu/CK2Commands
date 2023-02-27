package viviers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔巴拉罗曼AlbalaromaineBarony struct {
	feud.BaseBarony
}

var BAlbalaromaine阿尔巴拉罗曼 feud.Barony = &阿尔巴拉罗曼AlbalaromaineBarony{}

func init() {
    f := BAlbalaromaine阿尔巴拉罗曼.(*阿尔巴拉罗曼AlbalaromaineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albalaromaine",
		TitleName: "阿尔巴拉罗曼",
		TitleCode: "b_albalaromaine",
	}
}
