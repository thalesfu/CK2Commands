package uvek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌韦克UvekBarony struct {
	feud.BaseBarony
}

var BUvek乌韦克 feud.Barony = &乌韦克UvekBarony{}

func init() {
    f := BUvek乌韦克.(*乌韦克UvekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uvek",
		TitleName: "乌韦克",
		TitleCode: "b_uvek",
	}
}
