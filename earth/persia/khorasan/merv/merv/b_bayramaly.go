package merv

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜拉姆阿里BayramalyBarony struct {
	feud.BaseBarony
}

var BBayramaly拜拉姆阿里 feud.Barony = &拜拉姆阿里BayramalyBarony{}

func init() {
    f := BBayramaly拜拉姆阿里.(*拜拉姆阿里BayramalyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayramaly",
		TitleName: "拜拉姆阿里",
		TitleCode: "b_bayramaly",
	}
}
