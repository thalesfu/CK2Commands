package laksmanavati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗叉摩那伐底LaksmanavatiBarony struct {
	feud.BaseBarony
}

var BLaksmanavati罗叉摩那伐底 feud.Barony = &罗叉摩那伐底LaksmanavatiBarony{}

func init() {
	f := BLaksmanavati罗叉摩那伐底.(*罗叉摩那伐底LaksmanavatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laksmanavati",
		TitleName: "罗叉摩那伐底",
		TitleCode: "b_laksmanavati",
	}
}
