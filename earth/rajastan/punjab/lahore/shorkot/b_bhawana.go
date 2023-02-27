package shorkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴哈瓦那BhawanaBarony struct {
	feud.BaseBarony
}

var BBhawana巴哈瓦那 feud.Barony = &巴哈瓦那BhawanaBarony{}

func init() {
    f := BBhawana巴哈瓦那.(*巴哈瓦那BhawanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhawana",
		TitleName: "巴哈瓦那",
		TitleCode: "b_bhawana",
	}
}
