package padova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙塔尼亚纳MontagnanaBarony struct {
	feud.BaseBarony
}

var BMontagnana蒙塔尼亚纳 feud.Barony = &蒙塔尼亚纳MontagnanaBarony{}

func init() {
    f := BMontagnana蒙塔尼亚纳.(*蒙塔尼亚纳MontagnanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montagnana",
		TitleName: "蒙塔尼亚纳",
		TitleCode: "b_montagnana",
	}
}
