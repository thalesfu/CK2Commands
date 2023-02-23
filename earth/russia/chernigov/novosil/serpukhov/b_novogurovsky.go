package serpukhov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺沃古罗夫斯基NovogurovskyBarony struct {
	feud.BaseBarony
}

var BNovogurovsky诺沃古罗夫斯基 feud.Barony = &诺沃古罗夫斯基NovogurovskyBarony{}

func init() {
	f := BNovogurovsky诺沃古罗夫斯基.(*诺沃古罗夫斯基NovogurovskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novogurovsky",
		TitleName: "诺沃古罗夫斯基",
		TitleCode: "b_novogurovsky",
	}
}
