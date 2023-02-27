package anjou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕德LudeBarony struct {
	feud.BaseBarony
}

var BLude吕德 feud.Barony = &吕德LudeBarony{}

func init() {
    f := BLude吕德.(*吕德LudeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lude",
		TitleName: "吕德",
		TitleCode: "b_lude",
	}
}
