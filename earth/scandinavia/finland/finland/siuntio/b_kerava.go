package siuntio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯拉瓦KeravaBarony struct {
	feud.BaseBarony
}

var BKerava凯拉瓦 feud.Barony = &凯拉瓦KeravaBarony{}

func init() {
	f := BKerava凯拉瓦.(*凯拉瓦KeravaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerava",
		TitleName: "凯拉瓦",
		TitleCode: "b_kerava",
	}
}
