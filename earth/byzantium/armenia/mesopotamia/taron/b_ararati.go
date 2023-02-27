package taron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚拉拉特AraratiBarony struct {
	feud.BaseBarony
}

var BArarati亚拉拉特 feud.Barony = &亚拉拉特AraratiBarony{}

func init() {
    f := BArarati亚拉拉特.(*亚拉拉特AraratiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ararati",
		TitleName: "亚拉拉特",
		TitleCode: "b_ararati",
	}
}
