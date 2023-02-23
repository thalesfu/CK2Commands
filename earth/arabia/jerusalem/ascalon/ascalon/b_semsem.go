package ascalon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟姆森SemsemBarony struct {
	feud.BaseBarony
}

var BSemsem瑟姆森 feud.Barony = &瑟姆森SemsemBarony{}

func init() {
	f := BSemsem瑟姆森.(*瑟姆森SemsemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "semsem",
		TitleName: "瑟姆森",
		TitleCode: "b_semsem",
	}
}
