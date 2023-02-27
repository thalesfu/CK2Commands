package vikramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 浮卢婆BhurupalBarony struct {
	feud.BaseBarony
}

var BBhurupal浮卢婆 feud.Barony = &浮卢婆BhurupalBarony{}

func init() {
    f := BBhurupal浮卢婆.(*浮卢婆BhurupalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhurupal",
		TitleName: "浮卢婆",
		TitleCode: "b_bhurupal",
	}
}
