package euphrates

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯哈奇IshaqiBarony struct {
	feud.BaseBarony
}

var BIshaqi伊斯哈奇 feud.Barony = &伊斯哈奇IshaqiBarony{}

func init() {
    f := BIshaqi伊斯哈奇.(*伊斯哈奇IshaqiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ishaqi",
		TitleName: "伊斯哈奇",
		TitleCode: "b_ishaqi",
	}
}
