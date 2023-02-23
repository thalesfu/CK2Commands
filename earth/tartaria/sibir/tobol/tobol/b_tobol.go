package tobol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托博尔TobolBarony struct {
	feud.BaseBarony
}

var BTobol托博尔 feud.Barony = &托博尔TobolBarony{}

func init() {
	f := BTobol托博尔.(*托博尔TobolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tobol",
		TitleName: "托博尔",
		TitleCode: "b_tobol",
	}
}
