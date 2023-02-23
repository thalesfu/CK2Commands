package kyzyl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰丹ChadaanaBarony struct {
	feud.BaseBarony
}

var BChadaana恰丹 feud.Barony = &恰丹ChadaanaBarony{}

func init() {
	f := BChadaana恰丹.(*恰丹ChadaanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chadaana",
		TitleName: "恰丹",
		TitleCode: "b_chadaana",
	}
}
