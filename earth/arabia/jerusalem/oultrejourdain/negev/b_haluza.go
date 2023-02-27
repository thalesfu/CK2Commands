package negev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈鲁扎HaluzaBarony struct {
	feud.BaseBarony
}

var BHaluza哈鲁扎 feud.Barony = &哈鲁扎HaluzaBarony{}

func init() {
    f := BHaluza哈鲁扎.(*哈鲁扎HaluzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haluza",
		TitleName: "哈鲁扎",
		TitleCode: "b_haluza",
	}
}
