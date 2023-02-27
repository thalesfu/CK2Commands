package draa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朱阿堡KasbaeljouaBarony struct {
	feud.BaseBarony
}

var BKasbaeljoua朱阿堡 feud.Barony = &朱阿堡KasbaeljouaBarony{}

func init() {
    f := BKasbaeljoua朱阿堡.(*朱阿堡KasbaeljouaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kasbaeljoua",
		TitleName: "朱阿堡",
		TitleCode: "b_kasbaeljoua",
	}
}
