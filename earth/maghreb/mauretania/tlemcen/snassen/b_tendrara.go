package snassen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坦德拉拉TendraraBarony struct {
	feud.BaseBarony
}

var BTendrara坦德拉拉 feud.Barony = &坦德拉拉TendraraBarony{}

func init() {
    f := BTendrara坦德拉拉.(*坦德拉拉TendraraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tendrara",
		TitleName: "坦德拉拉",
		TitleCode: "b_tendrara",
	}
}
