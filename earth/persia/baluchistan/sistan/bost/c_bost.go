package bost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BostCounty interface {
	feud.County
	BBost博斯特() feud.Barony
	BChaman茶慢() feud.Barony
	BChora乔拉() feud.Barony
	BGereshk格里什克() feud.Barony
	BKandahar坎大哈() feud.Barony
	BPanjwai潘杰瓦伊() feud.Barony
	BSangin桑金() feud.Barony
}

type 博斯特BostCounty struct {
	feud.BaseCounty
	博斯特Bost     feud.Barony
	茶慢Chaman    feud.Barony
	乔拉Chora     feud.Barony
	格里什克Gereshk feud.Barony
	坎大哈Kandahar feud.Barony
	潘杰瓦伊Panjwai feud.Barony
	桑金Sangin    feud.Barony
}

func (c *博斯特BostCounty) BBost博斯特() feud.Barony {
	return c.博斯特Bost
}

func (c *博斯特BostCounty) BChaman茶慢() feud.Barony {
	return c.茶慢Chaman
}

func (c *博斯特BostCounty) BChora乔拉() feud.Barony {
	return c.乔拉Chora
}

func (c *博斯特BostCounty) BGereshk格里什克() feud.Barony {
	return c.格里什克Gereshk
}

func (c *博斯特BostCounty) BKandahar坎大哈() feud.Barony {
	return c.坎大哈Kandahar
}

func (c *博斯特BostCounty) BPanjwai潘杰瓦伊() feud.Barony {
	return c.潘杰瓦伊Panjwai
}

func (c *博斯特BostCounty) BSangin桑金() feud.Barony {
	return c.桑金Sangin
}

var CBost博斯特 BostCounty = &博斯特BostCounty{}

func init() {
	f := CBost博斯特.(*博斯特BostCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1183",
		Title:     "bost",
		TitleName: "博斯特",
		TitleCode: "c_bost",
		Baronies:  map[string]feud.Barony{},
	}

	f.博斯特Bost = BBost博斯特
	f.博斯特Bost.SetParent(f)

	f.茶慢Chaman = BChaman茶慢
	f.茶慢Chaman.SetParent(f)

	f.乔拉Chora = BChora乔拉
	f.乔拉Chora.SetParent(f)

	f.格里什克Gereshk = BGereshk格里什克
	f.格里什克Gereshk.SetParent(f)

	f.坎大哈Kandahar = BKandahar坎大哈
	f.坎大哈Kandahar.SetParent(f)

	f.潘杰瓦伊Panjwai = BPanjwai潘杰瓦伊
	f.潘杰瓦伊Panjwai.SetParent(f)

	f.桑金Sangin = BSangin桑金
	f.桑金Sangin.SetParent(f)

}
