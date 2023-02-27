package kremenchuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔拜AlbayBarony struct {
	feud.BaseBarony
}

var BAlbay阿尔拜 feud.Barony = &阿尔拜AlbayBarony{}

func init() {
    f := BAlbay阿尔拜.(*阿尔拜AlbayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albay",
		TitleName: "阿尔拜",
		TitleCode: "b_albay",
	}
}
