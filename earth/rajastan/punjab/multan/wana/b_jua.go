package wana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 周阿JuaBarony struct {
	feud.BaseBarony
}

var BJua周阿 feud.Barony = &周阿JuaBarony{}

func init() {
    f := BJua周阿.(*周阿JuaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jua",
		TitleName: "周阿",
		TitleCode: "b_jua",
	}
}
