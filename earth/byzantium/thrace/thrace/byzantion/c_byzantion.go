package byzantion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ByzantionCounty interface {
    feud.County
    BBlachernae弗拉赫尔奈() 	feud.Barony
    BConstantinople君士坦丁堡() 	feud.Barony
    BDeuteron狄特罗() 	feud.Barony
    BGalata加拉塔() 	feud.Barony
    BHagiasophia圣索菲亚() 	feud.Barony
    BHieron希埃隆() 	feud.Barony
    BPempton潘普顿() 	feud.Barony
    BVlanga乌兰戈() 	feud.Barony
}

type 君士坦丁堡ByzantionCounty struct {
	feud.BaseCounty
	弗拉赫尔奈Blachernae 	feud.Barony
	君士坦丁堡Constantinople 	feud.Barony
	狄特罗Deuteron 	feud.Barony
	加拉塔Galata 	feud.Barony
	圣索菲亚Hagiasophia 	feud.Barony
	希埃隆Hieron 	feud.Barony
	潘普顿Pempton 	feud.Barony
	乌兰戈Vlanga 	feud.Barony
}

func (c *君士坦丁堡ByzantionCounty) BBlachernae弗拉赫尔奈() feud.Barony {
	return c.弗拉赫尔奈Blachernae
}
    
func (c *君士坦丁堡ByzantionCounty) BConstantinople君士坦丁堡() feud.Barony {
	return c.君士坦丁堡Constantinople
}
    
func (c *君士坦丁堡ByzantionCounty) BDeuteron狄特罗() feud.Barony {
	return c.狄特罗Deuteron
}
    
func (c *君士坦丁堡ByzantionCounty) BGalata加拉塔() feud.Barony {
	return c.加拉塔Galata
}
    
func (c *君士坦丁堡ByzantionCounty) BHagiasophia圣索菲亚() feud.Barony {
	return c.圣索菲亚Hagiasophia
}
    
func (c *君士坦丁堡ByzantionCounty) BHieron希埃隆() feud.Barony {
	return c.希埃隆Hieron
}
    
func (c *君士坦丁堡ByzantionCounty) BPempton潘普顿() feud.Barony {
	return c.潘普顿Pempton
}
    
func (c *君士坦丁堡ByzantionCounty) BVlanga乌兰戈() feud.Barony {
	return c.乌兰戈Vlanga
}
    
var CByzantion君士坦丁堡 ByzantionCounty = &君士坦丁堡ByzantionCounty{}

func init() {
	f := CByzantion君士坦丁堡.(*君士坦丁堡ByzantionCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "496",
		Title:     "byzantion",
		TitleName: "君士坦丁堡",
		TitleCode: "c_byzantion",
		Baronies:  map[string]feud.Barony{},
	}

	f.弗拉赫尔奈Blachernae = BBlachernae弗拉赫尔奈
	f.弗拉赫尔奈Blachernae.SetParent(f)
	
	f.君士坦丁堡Constantinople = BConstantinople君士坦丁堡
	f.君士坦丁堡Constantinople.SetParent(f)
	
	f.狄特罗Deuteron = BDeuteron狄特罗
	f.狄特罗Deuteron.SetParent(f)
	
	f.加拉塔Galata = BGalata加拉塔
	f.加拉塔Galata.SetParent(f)
	
	f.圣索菲亚Hagiasophia = BHagiasophia圣索菲亚
	f.圣索菲亚Hagiasophia.SetParent(f)
	
	f.希埃隆Hieron = BHieron希埃隆
	f.希埃隆Hieron.SetParent(f)
	
	f.潘普顿Pempton = BPempton潘普顿
	f.潘普顿Pempton.SetParent(f)
	
	f.乌兰戈Vlanga = BVlanga乌兰戈
	f.乌兰戈Vlanga.SetParent(f)
	
}
