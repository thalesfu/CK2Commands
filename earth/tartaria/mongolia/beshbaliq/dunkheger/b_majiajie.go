package dunkheger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马家街MajiajieBarony struct {
	feud.BaseBarony
}

var BMajiajie马家街 feud.Barony = &马家街MajiajieBarony{}

func init() {
	f := BMajiajie马家街.(*马家街MajiajieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "majiajie",
		TitleName: "马家街",
		TitleCode: "b_majiajie",
	}
}
