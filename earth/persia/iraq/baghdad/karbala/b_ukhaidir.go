package karbala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌海迪尔UkhaidirBarony struct {
	feud.BaseBarony
}

var BUkhaidir乌海迪尔 feud.Barony = &乌海迪尔UkhaidirBarony{}

func init() {
	f := BUkhaidir乌海迪尔.(*乌海迪尔UkhaidirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ukhaidir",
		TitleName: "乌海迪尔",
		TitleCode: "b_ukhaidir",
	}
}
