package peremyshl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柳巴维奇LiubachivBarony struct {
	feud.BaseBarony
}

var BLiubachiv柳巴维奇 feud.Barony = &柳巴维奇LiubachivBarony{}

func init() {
    f := BLiubachiv柳巴维奇.(*柳巴维奇LiubachivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liubachiv",
		TitleName: "柳巴维奇",
		TitleCode: "b_liubachiv",
	}
}
