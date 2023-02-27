package damoh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 僧伽罗摩城SingrampurBarony struct {
	feud.BaseBarony
}

var BSingrampur僧伽罗摩城 feud.Barony = &僧伽罗摩城SingrampurBarony{}

func init() {
    f := BSingrampur僧伽罗摩城.(*僧伽罗摩城SingrampurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "singrampur",
		TitleName: "僧伽罗摩城",
		TitleCode: "b_singrampur",
	}
}
