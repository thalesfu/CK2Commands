package uda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兀的UdaBarony struct {
	feud.BaseBarony
}

var BUda兀的 feud.Barony = &兀的UdaBarony{}

func init() {
	f := BUda兀的.(*兀的UdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uda",
		TitleName: "兀的",
		TitleCode: "b_uda",
	}
}
