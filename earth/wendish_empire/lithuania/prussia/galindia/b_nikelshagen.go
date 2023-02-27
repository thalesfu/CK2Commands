package galindia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼克尔斯哈根NikelshagenBarony struct {
	feud.BaseBarony
}

var BNikelshagen尼克尔斯哈根 feud.Barony = &尼克尔斯哈根NikelshagenBarony{}

func init() {
    f := BNikelshagen尼克尔斯哈根.(*尼克尔斯哈根NikelshagenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nikelshagen",
		TitleName: "尼克尔斯哈根",
		TitleCode: "b_nikelshagen",
	}
}
