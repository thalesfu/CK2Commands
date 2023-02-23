package volkovysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日罗维奇ZyrovichyBarony struct {
	feud.BaseBarony
}

var BZyrovichy日罗维奇 feud.Barony = &日罗维奇ZyrovichyBarony{}

func init() {
	f := BZyrovichy日罗维奇.(*日罗维奇ZyrovichyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zyrovichy",
		TitleName: "日罗维奇",
		TitleCode: "b_zyrovichy",
	}
}
