package alania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特哈巴_叶尔德TkhabayerdyBarony struct {
	feud.BaseBarony
}

var BTkhabayerdy特哈巴_叶尔德 feud.Barony = &特哈巴_叶尔德TkhabayerdyBarony{}

func init() {
    f := BTkhabayerdy特哈巴_叶尔德.(*特哈巴_叶尔德TkhabayerdyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tkhabayerdy",
		TitleName: "特哈巴_叶尔德",
		TitleCode: "b_tkhabayerdy",
	}
}
