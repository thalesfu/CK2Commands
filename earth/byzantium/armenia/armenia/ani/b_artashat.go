package ani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔塔沙特ArtashatBarony struct {
	feud.BaseBarony
}

var BArtashat阿尔塔沙特 feud.Barony = &阿尔塔沙特ArtashatBarony{}

func init() {
	f := BArtashat阿尔塔沙特.(*阿尔塔沙特ArtashatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "artashat",
		TitleName: "阿尔塔沙特",
		TitleCode: "b_artashat",
	}
}
