package ascalon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AscalonCounty interface {
	feud.County
	BAgelen阿格伦() feud.Barony
	BAscalon亚实基伦() feud.Barony
	BBlanchegarde布兰彻贾德() feud.Barony
	BBothme博赫梅() feud.Barony
	BHarbijah赫比贾() feud.Barony
	BHuidre胡提尔() feud.Barony
	BLaforbie拉佛比() feud.Barony
	BSemsem瑟姆森() feud.Barony
}

type 亚实基伦AscalonCounty struct {
	feud.BaseCounty
	阿格伦Agelen         feud.Barony
	亚实基伦Ascalon       feud.Barony
	布兰彻贾德Blanchegarde feud.Barony
	博赫梅Bothme         feud.Barony
	赫比贾Harbijah       feud.Barony
	胡提尔Huidre         feud.Barony
	拉佛比Laforbie       feud.Barony
	瑟姆森Semsem         feud.Barony
}

func (c *亚实基伦AscalonCounty) BAgelen阿格伦() feud.Barony {
	return c.阿格伦Agelen
}

func (c *亚实基伦AscalonCounty) BAscalon亚实基伦() feud.Barony {
	return c.亚实基伦Ascalon
}

func (c *亚实基伦AscalonCounty) BBlanchegarde布兰彻贾德() feud.Barony {
	return c.布兰彻贾德Blanchegarde
}

func (c *亚实基伦AscalonCounty) BBothme博赫梅() feud.Barony {
	return c.博赫梅Bothme
}

func (c *亚实基伦AscalonCounty) BHarbijah赫比贾() feud.Barony {
	return c.赫比贾Harbijah
}

func (c *亚实基伦AscalonCounty) BHuidre胡提尔() feud.Barony {
	return c.胡提尔Huidre
}

func (c *亚实基伦AscalonCounty) BLaforbie拉佛比() feud.Barony {
	return c.拉佛比Laforbie
}

func (c *亚实基伦AscalonCounty) BSemsem瑟姆森() feud.Barony {
	return c.瑟姆森Semsem
}

var CAscalon亚实基伦 AscalonCounty = &亚实基伦AscalonCounty{}

func init() {
	f := CAscalon亚实基伦.(*亚实基伦AscalonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "780",
		Title:     "ascalon",
		TitleName: "亚实基伦",
		TitleCode: "c_ascalon",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格伦Agelen = BAgelen阿格伦
	f.阿格伦Agelen.SetParent(f)

	f.亚实基伦Ascalon = BAscalon亚实基伦
	f.亚实基伦Ascalon.SetParent(f)

	f.布兰彻贾德Blanchegarde = BBlanchegarde布兰彻贾德
	f.布兰彻贾德Blanchegarde.SetParent(f)

	f.博赫梅Bothme = BBothme博赫梅
	f.博赫梅Bothme.SetParent(f)

	f.赫比贾Harbijah = BHarbijah赫比贾
	f.赫比贾Harbijah.SetParent(f)

	f.胡提尔Huidre = BHuidre胡提尔
	f.胡提尔Huidre.SetParent(f)

	f.拉佛比Laforbie = BLaforbie拉佛比
	f.拉佛比Laforbie.SetParent(f)

	f.瑟姆森Semsem = BSemsem瑟姆森
	f.瑟姆森Semsem.SetParent(f)

}
