package bost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茶慢ChamanBarony struct {
	feud.BaseBarony
}

var BChaman茶慢 feud.Barony = &茶慢ChamanBarony{}

func init() {
	f := BChaman茶慢.(*茶慢ChamanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaman",
		TitleName: "茶慢",
		TitleCode: "b_chaman",
	}
}
