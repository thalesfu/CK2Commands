package kubera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杨桃山YangtaoshanBarony struct {
	feud.BaseBarony
}

var BYangtaoshan杨桃山 feud.Barony = &杨桃山YangtaoshanBarony{}

func init() {
	f := BYangtaoshan杨桃山.(*杨桃山YangtaoshanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yangtaoshan",
		TitleName: "杨桃山",
		TitleCode: "b_yangtaoshan",
	}
}
