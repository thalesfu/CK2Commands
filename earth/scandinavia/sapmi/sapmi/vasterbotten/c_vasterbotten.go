package vasterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VasterbottenCounty interface {
	feud.County
	BBygdea比格迪奥() feud.Barony
	BKalix卡利克斯() feud.Barony
	BLovanger勒翁厄尔() feud.Barony
	BLulea吕勒奥() feud.Barony
	BPitea皮特奥() feud.Barony
	BSkelleftea谢莱夫特奥() feud.Barony
	BTornea托尔尼奥() feud.Barony
	BUmea于默奥() feud.Barony
}

type 西博滕VasterbottenCounty struct {
	feud.BaseCounty
	比格迪奥Bygdea      feud.Barony
	卡利克斯Kalix       feud.Barony
	勒翁厄尔Lovanger    feud.Barony
	吕勒奥Lulea        feud.Barony
	皮特奥Pitea        feud.Barony
	谢莱夫特奥Skelleftea feud.Barony
	托尔尼奥Tornea      feud.Barony
	于默奥Umea         feud.Barony
}

func (c *西博滕VasterbottenCounty) BBygdea比格迪奥() feud.Barony {
	return c.比格迪奥Bygdea
}

func (c *西博滕VasterbottenCounty) BKalix卡利克斯() feud.Barony {
	return c.卡利克斯Kalix
}

func (c *西博滕VasterbottenCounty) BLovanger勒翁厄尔() feud.Barony {
	return c.勒翁厄尔Lovanger
}

func (c *西博滕VasterbottenCounty) BLulea吕勒奥() feud.Barony {
	return c.吕勒奥Lulea
}

func (c *西博滕VasterbottenCounty) BPitea皮特奥() feud.Barony {
	return c.皮特奥Pitea
}

func (c *西博滕VasterbottenCounty) BSkelleftea谢莱夫特奥() feud.Barony {
	return c.谢莱夫特奥Skelleftea
}

func (c *西博滕VasterbottenCounty) BTornea托尔尼奥() feud.Barony {
	return c.托尔尼奥Tornea
}

func (c *西博滕VasterbottenCounty) BUmea于默奥() feud.Barony {
	return c.于默奥Umea
}

var CVasterbotten西博滕 VasterbottenCounty = &西博滕VasterbottenCounty{}

func init() {
	f := CVasterbotten西博滕.(*西博滕VasterbottenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "280",
		Title:     "vasterbotten",
		TitleName: "西博滕",
		TitleCode: "c_vasterbotten",
		Baronies:  map[string]feud.Barony{},
	}

	f.比格迪奥Bygdea = BBygdea比格迪奥
	f.比格迪奥Bygdea.SetParent(f)

	f.卡利克斯Kalix = BKalix卡利克斯
	f.卡利克斯Kalix.SetParent(f)

	f.勒翁厄尔Lovanger = BLovanger勒翁厄尔
	f.勒翁厄尔Lovanger.SetParent(f)

	f.吕勒奥Lulea = BLulea吕勒奥
	f.吕勒奥Lulea.SetParent(f)

	f.皮特奥Pitea = BPitea皮特奥
	f.皮特奥Pitea.SetParent(f)

	f.谢莱夫特奥Skelleftea = BSkelleftea谢莱夫特奥
	f.谢莱夫特奥Skelleftea.SetParent(f)

	f.托尔尼奥Tornea = BTornea托尔尼奥
	f.托尔尼奥Tornea.SetParent(f)

	f.于默奥Umea = BUmea于默奥
	f.于默奥Umea.SetParent(f)

}
