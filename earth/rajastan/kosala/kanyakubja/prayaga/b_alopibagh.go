package prayaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卢毗婆迦AlopibaghBarony struct {
	feud.BaseBarony
}

var BAlopibagh阿卢毗婆迦 feud.Barony = &阿卢毗婆迦AlopibaghBarony{}

func init() {
	f := BAlopibagh阿卢毗婆迦.(*阿卢毗婆迦AlopibaghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alopibagh",
		TitleName: "阿卢毗婆迦",
		TitleCode: "b_alopibagh",
	}
}
