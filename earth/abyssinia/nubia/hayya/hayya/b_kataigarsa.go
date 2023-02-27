package hayya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塔伽萨KataigarsaBarony struct {
	feud.BaseBarony
}

var BKataigarsa卡塔伽萨 feud.Barony = &卡塔伽萨KataigarsaBarony{}

func init() {
    f := BKataigarsa卡塔伽萨.(*卡塔伽萨KataigarsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kataigarsa",
		TitleName: "卡塔伽萨",
		TitleCode: "b_kataigarsa",
	}
}
