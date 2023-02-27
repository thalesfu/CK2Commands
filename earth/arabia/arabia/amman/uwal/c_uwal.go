package uwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UwalCounty interface {
    feud.County
    BAldur杜尔() 	feud.Barony
    BHamala哈马拉() 	feud.Barony
    BJidda吉达() 	feud.Barony
    BManama麦纳麦() 	feud.Barony
    BMuharraq穆哈拉格() 	feud.Barony
    BSitra锡特拉() 	feud.Barony
    BUmm_al_nasan乌姆纳桑() 	feud.Barony
    BUmmalsabaan乌姆塞班() 	feud.Barony
}

type 乌瓦尔UwalCounty struct {
	feud.BaseCounty
	杜尔Aldur 	feud.Barony
	哈马拉Hamala 	feud.Barony
	吉达Jidda 	feud.Barony
	麦纳麦Manama 	feud.Barony
	穆哈拉格Muharraq 	feud.Barony
	锡特拉Sitra 	feud.Barony
	乌姆纳桑Umm_al_nasan 	feud.Barony
	乌姆塞班Ummalsabaan 	feud.Barony
}

func (c *乌瓦尔UwalCounty) BAldur杜尔() feud.Barony {
	return c.杜尔Aldur
}
    
func (c *乌瓦尔UwalCounty) BHamala哈马拉() feud.Barony {
	return c.哈马拉Hamala
}
    
func (c *乌瓦尔UwalCounty) BJidda吉达() feud.Barony {
	return c.吉达Jidda
}
    
func (c *乌瓦尔UwalCounty) BManama麦纳麦() feud.Barony {
	return c.麦纳麦Manama
}
    
func (c *乌瓦尔UwalCounty) BMuharraq穆哈拉格() feud.Barony {
	return c.穆哈拉格Muharraq
}
    
func (c *乌瓦尔UwalCounty) BSitra锡特拉() feud.Barony {
	return c.锡特拉Sitra
}
    
func (c *乌瓦尔UwalCounty) BUmm_al_nasan乌姆纳桑() feud.Barony {
	return c.乌姆纳桑Umm_al_nasan
}
    
func (c *乌瓦尔UwalCounty) BUmmalsabaan乌姆塞班() feud.Barony {
	return c.乌姆塞班Ummalsabaan
}
    
var CUwal乌瓦尔 UwalCounty = &乌瓦尔UwalCounty{}

func init() {
	f := CUwal乌瓦尔.(*乌瓦尔UwalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1531",
		Title:     "uwal",
		TitleName: "乌瓦尔",
		TitleCode: "c_uwal",
		Baronies:  map[string]feud.Barony{},
	}

	f.杜尔Aldur = BAldur杜尔
	f.杜尔Aldur.SetParent(f)
	
	f.哈马拉Hamala = BHamala哈马拉
	f.哈马拉Hamala.SetParent(f)
	
	f.吉达Jidda = BJidda吉达
	f.吉达Jidda.SetParent(f)
	
	f.麦纳麦Manama = BManama麦纳麦
	f.麦纳麦Manama.SetParent(f)
	
	f.穆哈拉格Muharraq = BMuharraq穆哈拉格
	f.穆哈拉格Muharraq.SetParent(f)
	
	f.锡特拉Sitra = BSitra锡特拉
	f.锡特拉Sitra.SetParent(f)
	
	f.乌姆纳桑Umm_al_nasan = BUmm_al_nasan乌姆纳桑
	f.乌姆纳桑Umm_al_nasan.SetParent(f)
	
	f.乌姆塞班Ummalsabaan = BUmmalsabaan乌姆塞班
	f.乌姆塞班Ummalsabaan.SetParent(f)
	
}
