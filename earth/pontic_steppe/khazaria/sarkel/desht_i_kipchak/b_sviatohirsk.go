package desht_i_kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯维亚托戈尔斯克SviatohirskBarony struct {
	feud.BaseBarony
}

var BSviatohirsk斯维亚托戈尔斯克 feud.Barony = &斯维亚托戈尔斯克SviatohirskBarony{}

func init() {
    f := BSviatohirsk斯维亚托戈尔斯克.(*斯维亚托戈尔斯克SviatohirskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sviatohirsk",
		TitleName: "斯维亚托戈尔斯克",
		TitleCode: "b_sviatohirsk",
	}
}
