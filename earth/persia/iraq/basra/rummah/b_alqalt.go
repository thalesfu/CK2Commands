package rummah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿喀拉特AlqaltBarony struct {
	feud.BaseBarony
}

var BAlqalt阿喀拉特 feud.Barony = &阿喀拉特AlqaltBarony{}

func init() {
	f := BAlqalt阿喀拉特.(*阿喀拉特AlqaltBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alqalt",
		TitleName: "阿喀拉特",
		TitleCode: "b_alqalt",
	}
}
