package belgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切塔泰亚阿尔巴Cetatea_albaBarony struct {
	feud.BaseBarony
}

var BCetatea_alba切塔泰亚阿尔巴 feud.Barony = &切塔泰亚阿尔巴Cetatea_albaBarony{}

func init() {
    f := BCetatea_alba切塔泰亚阿尔巴.(*切塔泰亚阿尔巴Cetatea_albaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cetatea_alba",
		TitleName: "切塔泰亚阿尔巴",
		TitleCode: "b_cetatea_alba",
	}
}
