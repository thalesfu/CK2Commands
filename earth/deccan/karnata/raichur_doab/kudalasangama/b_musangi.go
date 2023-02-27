package kudalasangama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆桑戈MusangiBarony struct {
	feud.BaseBarony
}

var BMusangi穆桑戈 feud.Barony = &穆桑戈MusangiBarony{}

func init() {
    f := BMusangi穆桑戈.(*穆桑戈MusangiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "musangi",
		TitleName: "穆桑戈",
		TitleCode: "b_musangi",
	}
}
