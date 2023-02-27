package oualata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔哈雷特TarhaletBarony struct {
	feud.BaseBarony
}

var BTarhalet塔尔哈雷特 feud.Barony = &塔尔哈雷特TarhaletBarony{}

func init() {
    f := BTarhalet塔尔哈雷特.(*塔尔哈雷特TarhaletBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarhalet",
		TitleName: "塔尔哈雷特",
		TitleCode: "b_tarhalet",
	}
}
