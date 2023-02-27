package ikh_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼勒德KhuldBarony struct {
	feud.BaseBarony
}

var BKhuld呼勒德 feud.Barony = &呼勒德KhuldBarony{}

func init() {
    f := BKhuld呼勒德.(*呼勒德KhuldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khuld",
		TitleName: "呼勒德",
		TitleCode: "b_khuld",
	}
}
