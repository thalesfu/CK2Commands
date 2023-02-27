package sedyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿基姆AkimBarony struct {
	feud.BaseBarony
}

var BAkim阿基姆 feud.Barony = &阿基姆AkimBarony{}

func init() {
    f := BAkim阿基姆.(*阿基姆AkimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akim",
		TitleName: "阿基姆",
		TitleCode: "b_akim",
	}
}
