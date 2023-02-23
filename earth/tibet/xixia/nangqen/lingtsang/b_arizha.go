package lingtsang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿日扎ArizhaBarony struct {
	feud.BaseBarony
}

var BArizha阿日扎 feud.Barony = &阿日扎ArizhaBarony{}

func init() {
	f := BArizha阿日扎.(*阿日扎ArizhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arizha",
		TitleName: "阿日扎",
		TitleCode: "b_arizha",
	}
}
