package como

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔比扎泰AlbizzateBarony struct {
	feud.BaseBarony
}

var BAlbizzate阿尔比扎泰 feud.Barony = &阿尔比扎泰AlbizzateBarony{}

func init() {
    f := BAlbizzate阿尔比扎泰.(*阿尔比扎泰AlbizzateBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albizzate",
		TitleName: "阿尔比扎泰",
		TitleCode: "b_albizzate",
	}
}
