package sjaelland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔丁堡VordingborgBarony struct {
	feud.BaseBarony
}

var BVordingborg沃尔丁堡 feud.Barony = &沃尔丁堡VordingborgBarony{}

func init() {
	f := BVordingborg沃尔丁堡.(*沃尔丁堡VordingborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vordingborg",
		TitleName: "沃尔丁堡",
		TitleCode: "b_vordingborg",
	}
}
