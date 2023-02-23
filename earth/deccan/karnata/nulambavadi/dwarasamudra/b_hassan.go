package dwarasamudra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈桑HassanBarony struct {
	feud.BaseBarony
}

var BHassan哈桑 feud.Barony = &哈桑HassanBarony{}

func init() {
	f := BHassan哈桑.(*哈桑HassanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hassan",
		TitleName: "哈桑",
		TitleCode: "b_hassan",
	}
}
