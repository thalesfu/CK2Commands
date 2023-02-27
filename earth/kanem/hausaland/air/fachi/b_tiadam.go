package fachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂阿达姆TiadamBarony struct {
	feud.BaseBarony
}

var BTiadam蒂阿达姆 feud.Barony = &蒂阿达姆TiadamBarony{}

func init() {
    f := BTiadam蒂阿达姆.(*蒂阿达姆TiadamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiadam",
		TitleName: "蒂阿达姆",
		TitleCode: "b_tiadam",
	}
}
