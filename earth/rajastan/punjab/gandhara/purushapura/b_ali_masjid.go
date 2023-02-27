package purushapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿里麦斯吉德Ali_masjidBarony struct {
	feud.BaseBarony
}

var BAli_masjid阿里麦斯吉德 feud.Barony = &阿里麦斯吉德Ali_masjidBarony{}

func init() {
    f := BAli_masjid阿里麦斯吉德.(*阿里麦斯吉德Ali_masjidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ali_masjid",
		TitleName: "阿里麦斯吉德",
		TitleCode: "b_ali_masjid",
	}
}
