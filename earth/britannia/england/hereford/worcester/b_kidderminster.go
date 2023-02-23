package worcester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基德明斯特KidderminsterBarony struct {
	feud.BaseBarony
}

var BKidderminster基德明斯特 feud.Barony = &基德明斯特KidderminsterBarony{}

func init() {
	f := BKidderminster基德明斯特.(*基德明斯特KidderminsterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kidderminster",
		TitleName: "基德明斯特",
		TitleCode: "b_kidderminster",
	}
}
