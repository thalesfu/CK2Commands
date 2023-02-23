package bharauli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BharauliCounty interface {
	feud.County
	BBhadri婆达利() feud.Barony
	BBharauli婆楼梨() feud.Barony
	BGotitoria瞿帝都梨耶() feud.Barony
	BHaktsarpur诃克娑补() feud.Barony
	BJevargi祗伐耆() feud.Barony
	BLalganj罗梨犍阇() feud.Barony
	BManikpur摩尼迦城() feud.Barony
}

type 婆楼梨BharauliCounty struct {
	feud.BaseCounty
	婆达利Bhadri      feud.Barony
	婆楼梨Bharauli    feud.Barony
	瞿帝都梨耶Gotitoria feud.Barony
	诃克娑补Haktsarpur feud.Barony
	祗伐耆Jevargi     feud.Barony
	罗梨犍阇Lalganj    feud.Barony
	摩尼迦城Manikpur   feud.Barony
}

func (c *婆楼梨BharauliCounty) BBhadri婆达利() feud.Barony {
	return c.婆达利Bhadri
}

func (c *婆楼梨BharauliCounty) BBharauli婆楼梨() feud.Barony {
	return c.婆楼梨Bharauli
}

func (c *婆楼梨BharauliCounty) BGotitoria瞿帝都梨耶() feud.Barony {
	return c.瞿帝都梨耶Gotitoria
}

func (c *婆楼梨BharauliCounty) BHaktsarpur诃克娑补() feud.Barony {
	return c.诃克娑补Haktsarpur
}

func (c *婆楼梨BharauliCounty) BJevargi祗伐耆() feud.Barony {
	return c.祗伐耆Jevargi
}

func (c *婆楼梨BharauliCounty) BLalganj罗梨犍阇() feud.Barony {
	return c.罗梨犍阇Lalganj
}

func (c *婆楼梨BharauliCounty) BManikpur摩尼迦城() feud.Barony {
	return c.摩尼迦城Manikpur
}

var CBharauli婆楼梨 BharauliCounty = &婆楼梨BharauliCounty{}

func init() {
	f := CBharauli婆楼梨.(*婆楼梨BharauliCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1284",
		Title:     "bharauli",
		TitleName: "婆楼梨",
		TitleCode: "c_bharauli",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆达利Bhadri = BBhadri婆达利
	f.婆达利Bhadri.SetParent(f)

	f.婆楼梨Bharauli = BBharauli婆楼梨
	f.婆楼梨Bharauli.SetParent(f)

	f.瞿帝都梨耶Gotitoria = BGotitoria瞿帝都梨耶
	f.瞿帝都梨耶Gotitoria.SetParent(f)

	f.诃克娑补Haktsarpur = BHaktsarpur诃克娑补
	f.诃克娑补Haktsarpur.SetParent(f)

	f.祗伐耆Jevargi = BJevargi祗伐耆
	f.祗伐耆Jevargi.SetParent(f)

	f.罗梨犍阇Lalganj = BLalganj罗梨犍阇
	f.罗梨犍阇Lalganj.SetParent(f)

	f.摩尼迦城Manikpur = BManikpur摩尼迦城
	f.摩尼迦城Manikpur.SetParent(f)

}
