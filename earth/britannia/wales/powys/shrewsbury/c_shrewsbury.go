package shrewsbury

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ShrewsburyCounty interface {
	feud.County
	BBridgnorth布里奇诺斯() feud.Barony
	BChirk彻克() feud.Barony
	BClun克兰() feud.Barony
	BLudlow拉德洛() feud.Barony
	BOswestry奥斯沃斯特里() feud.Barony
	BShrewsbury什鲁斯伯里() feud.Barony
	BWenlock文洛克() feud.Barony
	BWhitchurch惠特彻奇() feud.Barony
}

type 什鲁斯伯里ShrewsburyCounty struct {
	feud.BaseCounty
	布里奇诺斯Bridgnorth feud.Barony
	彻克Chirk         feud.Barony
	克兰Clun          feud.Barony
	拉德洛Ludlow       feud.Barony
	奥斯沃斯特里Oswestry  feud.Barony
	什鲁斯伯里Shrewsbury feud.Barony
	文洛克Wenlock      feud.Barony
	惠特彻奇Whitchurch  feud.Barony
}

func (c *什鲁斯伯里ShrewsburyCounty) BBridgnorth布里奇诺斯() feud.Barony {
	return c.布里奇诺斯Bridgnorth
}

func (c *什鲁斯伯里ShrewsburyCounty) BChirk彻克() feud.Barony {
	return c.彻克Chirk
}

func (c *什鲁斯伯里ShrewsburyCounty) BClun克兰() feud.Barony {
	return c.克兰Clun
}

func (c *什鲁斯伯里ShrewsburyCounty) BLudlow拉德洛() feud.Barony {
	return c.拉德洛Ludlow
}

func (c *什鲁斯伯里ShrewsburyCounty) BOswestry奥斯沃斯特里() feud.Barony {
	return c.奥斯沃斯特里Oswestry
}

func (c *什鲁斯伯里ShrewsburyCounty) BShrewsbury什鲁斯伯里() feud.Barony {
	return c.什鲁斯伯里Shrewsbury
}

func (c *什鲁斯伯里ShrewsburyCounty) BWenlock文洛克() feud.Barony {
	return c.文洛克Wenlock
}

func (c *什鲁斯伯里ShrewsburyCounty) BWhitchurch惠特彻奇() feud.Barony {
	return c.惠特彻奇Whitchurch
}

var CShrewsbury什鲁斯伯里 ShrewsburyCounty = &什鲁斯伯里ShrewsburyCounty{}

func init() {
	f := CShrewsbury什鲁斯伯里.(*什鲁斯伯里ShrewsburyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "66",
		Title:     "shrewsbury",
		TitleName: "什鲁斯伯里",
		TitleCode: "c_shrewsbury",
		Baronies:  map[string]feud.Barony{},
	}

	f.布里奇诺斯Bridgnorth = BBridgnorth布里奇诺斯
	f.布里奇诺斯Bridgnorth.SetParent(f)

	f.彻克Chirk = BChirk彻克
	f.彻克Chirk.SetParent(f)

	f.克兰Clun = BClun克兰
	f.克兰Clun.SetParent(f)

	f.拉德洛Ludlow = BLudlow拉德洛
	f.拉德洛Ludlow.SetParent(f)

	f.奥斯沃斯特里Oswestry = BOswestry奥斯沃斯特里
	f.奥斯沃斯特里Oswestry.SetParent(f)

	f.什鲁斯伯里Shrewsbury = BShrewsbury什鲁斯伯里
	f.什鲁斯伯里Shrewsbury.SetParent(f)

	f.文洛克Wenlock = BWenlock文洛克
	f.文洛克Wenlock.SetParent(f)

	f.惠特彻奇Whitchurch = BWhitchurch惠特彻奇
	f.惠特彻奇Whitchurch.SetParent(f)

}
