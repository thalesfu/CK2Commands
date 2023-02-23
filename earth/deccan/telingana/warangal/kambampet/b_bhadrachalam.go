package kambampet

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 跋陀罗遮蓝BhadrachalamBarony struct {
	feud.BaseBarony
}

var BBhadrachalam跋陀罗遮蓝 feud.Barony = &跋陀罗遮蓝BhadrachalamBarony{}

func init() {
	f := BBhadrachalam跋陀罗遮蓝.(*跋陀罗遮蓝BhadrachalamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhadrachalam",
		TitleName: "跋陀罗遮蓝",
		TitleCode: "b_bhadrachalam",
	}
}
