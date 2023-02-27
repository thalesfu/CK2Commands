package shetland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卡洛韦ScallowayBarony struct {
	feud.BaseBarony
}

var BScalloway斯卡洛韦 feud.Barony = &斯卡洛韦ScallowayBarony{}

func init() {
    f := BScalloway斯卡洛韦.(*斯卡洛韦ScallowayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "scalloway",
		TitleName: "斯卡洛韦",
		TitleCode: "b_scalloway",
	}
}
