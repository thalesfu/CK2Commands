package voin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮夏涅PishchaneBarony struct {
	feud.BaseBarony
}

var BPishchane皮夏涅 feud.Barony = &皮夏涅PishchaneBarony{}

func init() {
    f := BPishchane皮夏涅.(*皮夏涅PishchaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pishchane",
		TitleName: "皮夏涅",
		TitleCode: "b_pishchane",
	}
}
