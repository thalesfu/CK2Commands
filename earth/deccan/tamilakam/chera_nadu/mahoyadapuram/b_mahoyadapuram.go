package mahoyadapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩呼陀耶补罗MahoyadapuramBarony struct {
	feud.BaseBarony
}

var BMahoyadapuram摩呼陀耶补罗 feud.Barony = &摩呼陀耶补罗MahoyadapuramBarony{}

func init() {
    f := BMahoyadapuram摩呼陀耶补罗.(*摩呼陀耶补罗MahoyadapuramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahoyadapuram",
		TitleName: "摩呼陀耶补罗",
		TitleCode: "b_mahoyadapuram",
	}
}
