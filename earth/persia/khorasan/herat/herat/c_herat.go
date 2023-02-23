package herat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HeratCounty interface {
	feud.County
	BChishti契斯提() feud.Barony
	BFarsi法尔斯() feud.Barony
	BGulran古尔兰() feud.Barony
	BHerat亦鲁() feud.Barony
	BKarukh卡鲁赫() feud.Barony
	BKushk库什克() feud.Barony
	BObe奥比() feud.Barony
	BZarghun扎尔贡() feud.Barony
}

type 亦鲁HeratCounty struct {
	feud.BaseCounty
	契斯提Chishti feud.Barony
	法尔斯Farsi   feud.Barony
	古尔兰Gulran  feud.Barony
	亦鲁Herat    feud.Barony
	卡鲁赫Karukh  feud.Barony
	库什克Kushk   feud.Barony
	奥比Obe      feud.Barony
	扎尔贡Zarghun feud.Barony
}

func (c *亦鲁HeratCounty) BChishti契斯提() feud.Barony {
	return c.契斯提Chishti
}

func (c *亦鲁HeratCounty) BFarsi法尔斯() feud.Barony {
	return c.法尔斯Farsi
}

func (c *亦鲁HeratCounty) BGulran古尔兰() feud.Barony {
	return c.古尔兰Gulran
}

func (c *亦鲁HeratCounty) BHerat亦鲁() feud.Barony {
	return c.亦鲁Herat
}

func (c *亦鲁HeratCounty) BKarukh卡鲁赫() feud.Barony {
	return c.卡鲁赫Karukh
}

func (c *亦鲁HeratCounty) BKushk库什克() feud.Barony {
	return c.库什克Kushk
}

func (c *亦鲁HeratCounty) BObe奥比() feud.Barony {
	return c.奥比Obe
}

func (c *亦鲁HeratCounty) BZarghun扎尔贡() feud.Barony {
	return c.扎尔贡Zarghun
}

var CHerat亦鲁 HeratCounty = &亦鲁HeratCounty{}

func init() {
	f := CHerat亦鲁.(*亦鲁HeratCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "905",
		Title:     "herat",
		TitleName: "亦鲁",
		TitleCode: "c_herat",
		Baronies:  map[string]feud.Barony{},
	}

	f.契斯提Chishti = BChishti契斯提
	f.契斯提Chishti.SetParent(f)

	f.法尔斯Farsi = BFarsi法尔斯
	f.法尔斯Farsi.SetParent(f)

	f.古尔兰Gulran = BGulran古尔兰
	f.古尔兰Gulran.SetParent(f)

	f.亦鲁Herat = BHerat亦鲁
	f.亦鲁Herat.SetParent(f)

	f.卡鲁赫Karukh = BKarukh卡鲁赫
	f.卡鲁赫Karukh.SetParent(f)

	f.库什克Kushk = BKushk库什克
	f.库什克Kushk.SetParent(f)

	f.奥比Obe = BObe奥比
	f.奥比Obe.SetParent(f)

	f.扎尔贡Zarghun = BZarghun扎尔贡
	f.扎尔贡Zarghun.SetParent(f)

}
