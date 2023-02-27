package kotivarsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯梨古利SiliguriBarony struct {
	feud.BaseBarony
}

var BSiliguri斯梨古利 feud.Barony = &斯梨古利SiliguriBarony{}

func init() {
    f := BSiliguri斯梨古利.(*斯梨古利SiliguriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siliguri",
		TitleName: "斯梨古利",
		TitleCode: "b_siliguri",
	}
}
