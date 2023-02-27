package beshbaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柳湾乡LiuwanxiangBarony struct {
	feud.BaseBarony
}

var BLiuwanxiang柳湾乡 feud.Barony = &柳湾乡LiuwanxiangBarony{}

func init() {
    f := BLiuwanxiang柳湾乡.(*柳湾乡LiuwanxiangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liuwanxiang",
		TitleName: "柳湾乡",
		TitleCode: "b_liuwanxiang",
	}
}
