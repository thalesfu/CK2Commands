package lincoln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格兰瑟姆GranthamBarony struct {
	feud.BaseBarony
}

var BGrantham格兰瑟姆 feud.Barony = &格兰瑟姆GranthamBarony{}

func init() {
    f := BGrantham格兰瑟姆.(*格兰瑟姆GranthamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grantham",
		TitleName: "格兰瑟姆",
		TitleCode: "b_grantham",
	}
}
