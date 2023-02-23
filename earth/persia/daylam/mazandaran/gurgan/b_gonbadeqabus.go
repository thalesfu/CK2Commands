package gurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡巴德卡武斯GonbadeqabusBarony struct {
	feud.BaseBarony
}

var BGonbadeqabus贡巴德卡武斯 feud.Barony = &贡巴德卡武斯GonbadeqabusBarony{}

func init() {
	f := BGonbadeqabus贡巴德卡武斯.(*贡巴德卡武斯GonbadeqabusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gonbadeqabus",
		TitleName: "贡巴德卡武斯",
		TitleCode: "b_gonbadeqabus",
	}
}
