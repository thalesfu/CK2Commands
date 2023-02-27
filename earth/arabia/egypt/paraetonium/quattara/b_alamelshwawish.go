package quattara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉姆沙锡AlamelshwawishBarony struct {
	feud.BaseBarony
}

var BAlamelshwawish阿拉姆沙锡 feud.Barony = &阿拉姆沙锡AlamelshwawishBarony{}

func init() {
    f := BAlamelshwawish阿拉姆沙锡.(*阿拉姆沙锡AlamelshwawishBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alamelshwawish",
		TitleName: "阿拉姆沙锡",
		TitleCode: "b_alamelshwawish",
	}
}
