package bahrein

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡韦尔Al_khuwayrBarony struct {
	feud.BaseBarony
}

var BAl_khuwayr胡韦尔 feud.Barony = &胡韦尔Al_khuwayrBarony{}

func init() {
    f := BAl_khuwayr胡韦尔.(*胡韦尔Al_khuwayrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_khuwayr",
		TitleName: "胡韦尔",
		TitleCode: "b_al_khuwayr",
	}
}
