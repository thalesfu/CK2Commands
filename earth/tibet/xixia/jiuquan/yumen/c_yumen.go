package yumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YumenCounty interface {
	feud.County
	BBaoen报恩() feud.Barony
	BChuanbei川北() feud.Barony
	BDongshuwo东树窝() feud.Barony
	BHuangzhawan黄闸湾() feud.Barony
	BShagang沙岗() feud.Barony
	BXiaxihao下西号() feud.Barony
	BYumen玉门() feud.Barony
}

type 玉门YumenCounty struct {
	feud.BaseCounty
	报恩Baoen        feud.Barony
	川北Chuanbei     feud.Barony
	东树窝Dongshuwo   feud.Barony
	黄闸湾Huangzhawan feud.Barony
	沙岗Shagang      feud.Barony
	下西号Xiaxihao    feud.Barony
	玉门Yumen        feud.Barony
}

func (c *玉门YumenCounty) BBaoen报恩() feud.Barony {
	return c.报恩Baoen
}

func (c *玉门YumenCounty) BChuanbei川北() feud.Barony {
	return c.川北Chuanbei
}

func (c *玉门YumenCounty) BDongshuwo东树窝() feud.Barony {
	return c.东树窝Dongshuwo
}

func (c *玉门YumenCounty) BHuangzhawan黄闸湾() feud.Barony {
	return c.黄闸湾Huangzhawan
}

func (c *玉门YumenCounty) BShagang沙岗() feud.Barony {
	return c.沙岗Shagang
}

func (c *玉门YumenCounty) BXiaxihao下西号() feud.Barony {
	return c.下西号Xiaxihao
}

func (c *玉门YumenCounty) BYumen玉门() feud.Barony {
	return c.玉门Yumen
}

var CYumen玉门 YumenCounty = &玉门YumenCounty{}

func init() {
	f := CYumen玉门.(*玉门YumenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1568",
		Title:     "yumen",
		TitleName: "玉门",
		TitleCode: "c_yumen",
		Baronies:  map[string]feud.Barony{},
	}

	f.报恩Baoen = BBaoen报恩
	f.报恩Baoen.SetParent(f)

	f.川北Chuanbei = BChuanbei川北
	f.川北Chuanbei.SetParent(f)

	f.东树窝Dongshuwo = BDongshuwo东树窝
	f.东树窝Dongshuwo.SetParent(f)

	f.黄闸湾Huangzhawan = BHuangzhawan黄闸湾
	f.黄闸湾Huangzhawan.SetParent(f)

	f.沙岗Shagang = BShagang沙岗
	f.沙岗Shagang.SetParent(f)

	f.下西号Xiaxihao = BXiaxihao下西号
	f.下西号Xiaxihao.SetParent(f)

	f.玉门Yumen = BYumen玉门
	f.玉门Yumen.SetParent(f)

}
