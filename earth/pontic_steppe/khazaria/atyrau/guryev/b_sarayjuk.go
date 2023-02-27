package guryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 小萨赖SarayjukBarony struct {
	feud.BaseBarony
}

var BSarayjuk小萨赖 feud.Barony = &小萨赖SarayjukBarony{}

func init() {
    f := BSarayjuk小萨赖.(*小萨赖SarayjukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarayjuk",
		TitleName: "小萨赖",
		TitleCode: "b_sarayjuk",
	}
}
