package yuryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加夫里洛夫镇Gavrilov_posadBarony struct {
	feud.BaseBarony
}

var BGavrilov_posad加夫里洛夫镇 feud.Barony = &加夫里洛夫镇Gavrilov_posadBarony{}

func init() {
    f := BGavrilov_posad加夫里洛夫镇.(*加夫里洛夫镇Gavrilov_posadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gavrilov_posad",
		TitleName: "加夫里洛夫镇",
		TitleCode: "b_gavrilov_posad",
	}
}
