package dakhina_desa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗依伽摩RayigamaBarony struct {
	feud.BaseBarony
}

var BRayigama罗依伽摩 feud.Barony = &罗依伽摩RayigamaBarony{}

func init() {
    f := BRayigama罗依伽摩.(*罗依伽摩RayigamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rayigama",
		TitleName: "罗依伽摩",
		TitleCode: "b_rayigama",
	}
}
