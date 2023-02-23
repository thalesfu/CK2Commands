package palmyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉哈尔JiharBarony struct {
	feud.BaseBarony
}

var BJihar吉哈尔 feud.Barony = &吉哈尔JiharBarony{}

func init() {
	f := BJihar吉哈尔.(*吉哈尔JiharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jihar",
		TitleName: "吉哈尔",
		TitleCode: "b_jihar",
	}
}
