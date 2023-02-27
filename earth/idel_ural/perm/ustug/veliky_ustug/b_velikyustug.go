package veliky_ustug

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大乌斯秋格VelikyustugBarony struct {
	feud.BaseBarony
}

var BVelikyustug大乌斯秋格 feud.Barony = &大乌斯秋格VelikyustugBarony{}

func init() {
    f := BVelikyustug大乌斯秋格.(*大乌斯秋格VelikyustugBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velikyustug",
		TitleName: "大乌斯秋格",
		TitleCode: "b_velikyustug",
	}
}
