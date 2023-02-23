package vaspurakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VaspurakanCounty interface {
	feud.County
	BAghtamar阿赫塔马尔() feud.Barony
	BBakear巴克尔() feud.Barony
	BHadamakert哈达马克特() feud.Barony
	BHaykaberd哈卡贝德() feud.Barony
	BSurbkhach苏尔布哈奇() feud.Barony
	BVan凡() feud.Barony
	BVaragavank瓦拉格凡克() feud.Barony
	BVostan沃斯坦() feud.Barony
}

type 瓦斯普拉坎VaspurakanCounty struct {
	feud.BaseCounty
	阿赫塔马尔Aghtamar   feud.Barony
	巴克尔Bakear       feud.Barony
	哈达马克特Hadamakert feud.Barony
	哈卡贝德Haykaberd   feud.Barony
	苏尔布哈奇Surbkhach  feud.Barony
	凡Van            feud.Barony
	瓦拉格凡克Varagavank feud.Barony
	沃斯坦Vostan       feud.Barony
}

func (c *瓦斯普拉坎VaspurakanCounty) BAghtamar阿赫塔马尔() feud.Barony {
	return c.阿赫塔马尔Aghtamar
}

func (c *瓦斯普拉坎VaspurakanCounty) BBakear巴克尔() feud.Barony {
	return c.巴克尔Bakear
}

func (c *瓦斯普拉坎VaspurakanCounty) BHadamakert哈达马克特() feud.Barony {
	return c.哈达马克特Hadamakert
}

func (c *瓦斯普拉坎VaspurakanCounty) BHaykaberd哈卡贝德() feud.Barony {
	return c.哈卡贝德Haykaberd
}

func (c *瓦斯普拉坎VaspurakanCounty) BSurbkhach苏尔布哈奇() feud.Barony {
	return c.苏尔布哈奇Surbkhach
}

func (c *瓦斯普拉坎VaspurakanCounty) BVan凡() feud.Barony {
	return c.凡Van
}

func (c *瓦斯普拉坎VaspurakanCounty) BVaragavank瓦拉格凡克() feud.Barony {
	return c.瓦拉格凡克Varagavank
}

func (c *瓦斯普拉坎VaspurakanCounty) BVostan沃斯坦() feud.Barony {
	return c.沃斯坦Vostan
}

var CVaspurakan瓦斯普拉坎 VaspurakanCounty = &瓦斯普拉坎VaspurakanCounty{}

func init() {
	f := CVaspurakan瓦斯普拉坎.(*瓦斯普拉坎VaspurakanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "682",
		Title:     "vaspurakan",
		TitleName: "瓦斯普拉坎",
		TitleCode: "c_vaspurakan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赫塔马尔Aghtamar = BAghtamar阿赫塔马尔
	f.阿赫塔马尔Aghtamar.SetParent(f)

	f.巴克尔Bakear = BBakear巴克尔
	f.巴克尔Bakear.SetParent(f)

	f.哈达马克特Hadamakert = BHadamakert哈达马克特
	f.哈达马克特Hadamakert.SetParent(f)

	f.哈卡贝德Haykaberd = BHaykaberd哈卡贝德
	f.哈卡贝德Haykaberd.SetParent(f)

	f.苏尔布哈奇Surbkhach = BSurbkhach苏尔布哈奇
	f.苏尔布哈奇Surbkhach.SetParent(f)

	f.凡Van = BVan凡
	f.凡Van.SetParent(f)

	f.瓦拉格凡克Varagavank = BVaragavank瓦拉格凡克
	f.瓦拉格凡克Varagavank.SetParent(f)

	f.沃斯坦Vostan = BVostan沃斯坦
	f.沃斯坦Vostan.SetParent(f)

}
