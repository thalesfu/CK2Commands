package laksmanavati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩罗陀MaldaBarony struct {
	feud.BaseBarony
}

var BMalda摩罗陀 feud.Barony = &摩罗陀MaldaBarony{}

func init() {
    f := BMalda摩罗陀.(*摩罗陀MaldaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malda",
		TitleName: "摩罗陀",
		TitleCode: "b_malda",
	}
}
