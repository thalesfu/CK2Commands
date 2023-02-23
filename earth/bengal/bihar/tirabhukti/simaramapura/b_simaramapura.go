package simaramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 僧摩罗摩城SimaramapuraBarony struct {
	feud.BaseBarony
}

var BSimaramapura僧摩罗摩城 feud.Barony = &僧摩罗摩城SimaramapuraBarony{}

func init() {
	f := BSimaramapura僧摩罗摩城.(*僧摩罗摩城SimaramapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "simaramapura",
		TitleName: "僧摩罗摩城",
		TitleCode: "b_simaramapura",
	}
}
