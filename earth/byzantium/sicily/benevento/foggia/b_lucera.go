package foggia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢切拉LuceraBarony struct {
	feud.BaseBarony
}

var BLucera卢切拉 feud.Barony = &卢切拉LuceraBarony{}

func init() {
    f := BLucera卢切拉.(*卢切拉LuceraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lucera",
		TitleName: "卢切拉",
		TitleCode: "b_lucera",
	}
}
