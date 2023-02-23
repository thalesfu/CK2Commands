package albania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 楚霍尔卡巴拉ChukhurkabalaBarony struct {
	feud.BaseBarony
}

var BChukhurkabala楚霍尔卡巴拉 feud.Barony = &楚霍尔卡巴拉ChukhurkabalaBarony{}

func init() {
	f := BChukhurkabala楚霍尔卡巴拉.(*楚霍尔卡巴拉ChukhurkabalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chukhurkabala",
		TitleName: "楚霍尔卡巴拉",
		TitleCode: "b_chukhurkabala",
	}
}
