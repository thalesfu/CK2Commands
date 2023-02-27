package verden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基希瓦尔瑟德KirchwalsedeBarony struct {
	feud.BaseBarony
}

var BKirchwalsede基希瓦尔瑟德 feud.Barony = &基希瓦尔瑟德KirchwalsedeBarony{}

func init() {
    f := BKirchwalsede基希瓦尔瑟德.(*基希瓦尔瑟德KirchwalsedeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirchwalsede",
		TitleName: "基希瓦尔瑟德",
		TitleCode: "b_kirchwalsede",
	}
}
