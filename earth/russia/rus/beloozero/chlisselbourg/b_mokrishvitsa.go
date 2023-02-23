package chlisselbourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫克里什维察MokrishvitsaBarony struct {
	feud.BaseBarony
}

var BMokrishvitsa莫克里什维察 feud.Barony = &莫克里什维察MokrishvitsaBarony{}

func init() {
	f := BMokrishvitsa莫克里什维察.(*莫克里什维察MokrishvitsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mokrishvitsa",
		TitleName: "莫克里什维察",
		TitleCode: "b_mokrishvitsa",
	}
}
