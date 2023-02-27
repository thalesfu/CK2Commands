package vaspurakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿赫塔马尔AghtamarBarony struct {
	feud.BaseBarony
}

var BAghtamar阿赫塔马尔 feud.Barony = &阿赫塔马尔AghtamarBarony{}

func init() {
    f := BAghtamar阿赫塔马尔.(*阿赫塔马尔AghtamarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aghtamar",
		TitleName: "阿赫塔马尔",
		TitleCode: "b_aghtamar",
	}
}
