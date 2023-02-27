package capua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡普阿CapuaBarony struct {
	feud.BaseBarony
}

var BCapua卡普阿 feud.Barony = &卡普阿CapuaBarony{}

func init() {
    f := BCapua卡普阿.(*卡普阿CapuaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "capua",
		TitleName: "卡普阿",
		TitleCode: "b_capua",
	}
}
