package lolland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳克斯考NacascoghBarony struct {
	feud.BaseBarony
}

var BNacascogh纳克斯考 feud.Barony = &纳克斯考NacascoghBarony{}

func init() {
	f := BNacascogh纳克斯考.(*纳克斯考NacascoghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nacascogh",
		TitleName: "纳克斯考",
		TitleCode: "b_nacascogh",
	}
}
