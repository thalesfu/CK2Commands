package gurjaratra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GurjaratraCounty interface {
	feud.County
	BBhaddar婆陀() feud.Barony
	BGurjaratra瞿折罗多罗() feud.Barony
	BHakla哈卡拉() feud.Barony
	BJhelum杰赫勒姆() feud.Barony
	BKarianwala加利安瓦拉() feud.Barony
	BKharian贾里安() feud.Barony
	BKunjah贡迦() feud.Barony
	BTibbi提毗() feud.Barony
}

type 瞿折罗多罗GurjaratraCounty struct {
	feud.BaseCounty
	婆陀Bhaddar       feud.Barony
	瞿折罗多罗Gurjaratra feud.Barony
	哈卡拉Hakla        feud.Barony
	杰赫勒姆Jhelum      feud.Barony
	加利安瓦拉Karianwala feud.Barony
	贾里安Kharian      feud.Barony
	贡迦Kunjah        feud.Barony
	提毗Tibbi         feud.Barony
}

func (c *瞿折罗多罗GurjaratraCounty) BBhaddar婆陀() feud.Barony {
	return c.婆陀Bhaddar
}

func (c *瞿折罗多罗GurjaratraCounty) BGurjaratra瞿折罗多罗() feud.Barony {
	return c.瞿折罗多罗Gurjaratra
}

func (c *瞿折罗多罗GurjaratraCounty) BHakla哈卡拉() feud.Barony {
	return c.哈卡拉Hakla
}

func (c *瞿折罗多罗GurjaratraCounty) BJhelum杰赫勒姆() feud.Barony {
	return c.杰赫勒姆Jhelum
}

func (c *瞿折罗多罗GurjaratraCounty) BKarianwala加利安瓦拉() feud.Barony {
	return c.加利安瓦拉Karianwala
}

func (c *瞿折罗多罗GurjaratraCounty) BKharian贾里安() feud.Barony {
	return c.贾里安Kharian
}

func (c *瞿折罗多罗GurjaratraCounty) BKunjah贡迦() feud.Barony {
	return c.贡迦Kunjah
}

func (c *瞿折罗多罗GurjaratraCounty) BTibbi提毗() feud.Barony {
	return c.提毗Tibbi
}

var CGurjaratra瞿折罗多罗 GurjaratraCounty = &瞿折罗多罗GurjaratraCounty{}

func init() {
	f := CGurjaratra瞿折罗多罗.(*瞿折罗多罗GurjaratraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1190",
		Title:     "gurjaratra",
		TitleName: "瞿折罗多罗",
		TitleCode: "c_gurjaratra",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆陀Bhaddar = BBhaddar婆陀
	f.婆陀Bhaddar.SetParent(f)

	f.瞿折罗多罗Gurjaratra = BGurjaratra瞿折罗多罗
	f.瞿折罗多罗Gurjaratra.SetParent(f)

	f.哈卡拉Hakla = BHakla哈卡拉
	f.哈卡拉Hakla.SetParent(f)

	f.杰赫勒姆Jhelum = BJhelum杰赫勒姆
	f.杰赫勒姆Jhelum.SetParent(f)

	f.加利安瓦拉Karianwala = BKarianwala加利安瓦拉
	f.加利安瓦拉Karianwala.SetParent(f)

	f.贾里安Kharian = BKharian贾里安
	f.贾里安Kharian.SetParent(f)

	f.贡迦Kunjah = BKunjah贡迦
	f.贡迦Kunjah.SetParent(f)

	f.提毗Tibbi = BTibbi提毗
	f.提毗Tibbi.SetParent(f)

}
