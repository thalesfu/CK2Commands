package loon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦森贝格WassenbergBarony struct {
	feud.BaseBarony
}

var BWassenberg瓦森贝格 feud.Barony = &瓦森贝格WassenbergBarony{}

func init() {
    f := BWassenberg瓦森贝格.(*瓦森贝格WassenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wassenberg",
		TitleName: "瓦森贝格",
		TitleCode: "b_wassenberg",
	}
}
