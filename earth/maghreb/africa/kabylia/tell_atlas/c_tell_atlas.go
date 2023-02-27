package tell_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Tell_atlasCounty interface {
    feud.County
    BDraaelmizan德拉米赞() 	feud.Barony
    BIkjan伊克詹() 	feud.Barony
    BMechreck梅彻尔克() 	feud.Barony
    BNacina纳斯那() 	feud.Barony
    BSetif塞提夫() 	feud.Barony
    BTadjenanet塔杰纳奈特() 	feud.Barony
    BTagoughalt塔古格哈() 	feud.Barony
    BZaoula祖拉() 	feud.Barony
}

type 塞提夫Tell_atlasCounty struct {
	feud.BaseCounty
	德拉米赞Draaelmizan 	feud.Barony
	伊克詹Ikjan 	feud.Barony
	梅彻尔克Mechreck 	feud.Barony
	纳斯那Nacina 	feud.Barony
	塞提夫Setif 	feud.Barony
	塔杰纳奈特Tadjenanet 	feud.Barony
	塔古格哈Tagoughalt 	feud.Barony
	祖拉Zaoula 	feud.Barony
}

func (c *塞提夫Tell_atlasCounty) BDraaelmizan德拉米赞() feud.Barony {
	return c.德拉米赞Draaelmizan
}
    
func (c *塞提夫Tell_atlasCounty) BIkjan伊克詹() feud.Barony {
	return c.伊克詹Ikjan
}
    
func (c *塞提夫Tell_atlasCounty) BMechreck梅彻尔克() feud.Barony {
	return c.梅彻尔克Mechreck
}
    
func (c *塞提夫Tell_atlasCounty) BNacina纳斯那() feud.Barony {
	return c.纳斯那Nacina
}
    
func (c *塞提夫Tell_atlasCounty) BSetif塞提夫() feud.Barony {
	return c.塞提夫Setif
}
    
func (c *塞提夫Tell_atlasCounty) BTadjenanet塔杰纳奈特() feud.Barony {
	return c.塔杰纳奈特Tadjenanet
}
    
func (c *塞提夫Tell_atlasCounty) BTagoughalt塔古格哈() feud.Barony {
	return c.塔古格哈Tagoughalt
}
    
func (c *塞提夫Tell_atlasCounty) BZaoula祖拉() feud.Barony {
	return c.祖拉Zaoula
}
    
var CTell_atlas塞提夫 Tell_atlasCounty = &塞提夫Tell_atlasCounty{}

func init() {
	f := CTell_atlas塞提夫.(*塞提夫Tell_atlasCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "824",
		Title:     "tell_atlas",
		TitleName: "塞提夫",
		TitleCode: "c_tell_atlas",
		Baronies:  map[string]feud.Barony{},
	}

	f.德拉米赞Draaelmizan = BDraaelmizan德拉米赞
	f.德拉米赞Draaelmizan.SetParent(f)
	
	f.伊克詹Ikjan = BIkjan伊克詹
	f.伊克詹Ikjan.SetParent(f)
	
	f.梅彻尔克Mechreck = BMechreck梅彻尔克
	f.梅彻尔克Mechreck.SetParent(f)
	
	f.纳斯那Nacina = BNacina纳斯那
	f.纳斯那Nacina.SetParent(f)
	
	f.塞提夫Setif = BSetif塞提夫
	f.塞提夫Setif.SetParent(f)
	
	f.塔杰纳奈特Tadjenanet = BTadjenanet塔杰纳奈特
	f.塔杰纳奈特Tadjenanet.SetParent(f)
	
	f.塔古格哈Tagoughalt = BTagoughalt塔古格哈
	f.塔古格哈Tagoughalt.SetParent(f)
	
	f.祖拉Zaoula = BZaoula祖拉
	f.祖拉Zaoula.SetParent(f)
	
}
