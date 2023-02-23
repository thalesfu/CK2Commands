package lopnor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LopnorCounty interface {
	feud.County
	BBeidijiadun北地家墩() feud.Barony
	BBibian壁边() feud.Barony
	BDingjiazhai丁家寨() feud.Barony
	BLopnor罗布泊() feud.Barony
	BNanhedong南河东() feud.Barony
	BNantung南屯() feud.Barony
	BQitun七屯() feud.Barony
}

type 罗布泊LopnorCounty struct {
	feud.BaseCounty
	北地家墩Beidijiadun feud.Barony
	壁边Bibian        feud.Barony
	丁家寨Dingjiazhai  feud.Barony
	罗布泊Lopnor       feud.Barony
	南河东Nanhedong    feud.Barony
	南屯Nantung       feud.Barony
	七屯Qitun         feud.Barony
}

func (c *罗布泊LopnorCounty) BBeidijiadun北地家墩() feud.Barony {
	return c.北地家墩Beidijiadun
}

func (c *罗布泊LopnorCounty) BBibian壁边() feud.Barony {
	return c.壁边Bibian
}

func (c *罗布泊LopnorCounty) BDingjiazhai丁家寨() feud.Barony {
	return c.丁家寨Dingjiazhai
}

func (c *罗布泊LopnorCounty) BLopnor罗布泊() feud.Barony {
	return c.罗布泊Lopnor
}

func (c *罗布泊LopnorCounty) BNanhedong南河东() feud.Barony {
	return c.南河东Nanhedong
}

func (c *罗布泊LopnorCounty) BNantung南屯() feud.Barony {
	return c.南屯Nantung
}

func (c *罗布泊LopnorCounty) BQitun七屯() feud.Barony {
	return c.七屯Qitun
}

var CLopnor罗布泊 LopnorCounty = &罗布泊LopnorCounty{}

func init() {
	f := CLopnor罗布泊.(*罗布泊LopnorCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1519",
		Title:     "lopnor",
		TitleName: "罗布泊",
		TitleCode: "c_lopnor",
		Baronies:  map[string]feud.Barony{},
	}

	f.北地家墩Beidijiadun = BBeidijiadun北地家墩
	f.北地家墩Beidijiadun.SetParent(f)

	f.壁边Bibian = BBibian壁边
	f.壁边Bibian.SetParent(f)

	f.丁家寨Dingjiazhai = BDingjiazhai丁家寨
	f.丁家寨Dingjiazhai.SetParent(f)

	f.罗布泊Lopnor = BLopnor罗布泊
	f.罗布泊Lopnor.SetParent(f)

	f.南河东Nanhedong = BNanhedong南河东
	f.南河东Nanhedong.SetParent(f)

	f.南屯Nantung = BNantung南屯
	f.南屯Nantung.SetParent(f)

	f.七屯Qitun = BQitun七屯
	f.七屯Qitun.SetParent(f)

}
