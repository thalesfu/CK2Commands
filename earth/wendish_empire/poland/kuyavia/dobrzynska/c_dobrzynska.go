package dobrzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DobrzynskaCounty interface {
    feud.County
    BDobrzyn多布任() 	feud.Barony
    BKikot基库乌() 	feud.Barony
    BLipno利普诺() 	feud.Barony
    BMichalowo米哈沃沃() 	feud.Barony
    BRypin雷平() 	feud.Barony
    BSkepe斯肯佩() 	feud.Barony
    BWloclawek奥斯特罗维泰() 	feud.Barony
}

type 多布任DobrzynskaCounty struct {
	feud.BaseCounty
	多布任Dobrzyn 	feud.Barony
	基库乌Kikot 	feud.Barony
	利普诺Lipno 	feud.Barony
	米哈沃沃Michalowo 	feud.Barony
	雷平Rypin 	feud.Barony
	斯肯佩Skepe 	feud.Barony
	奥斯特罗维泰Wloclawek 	feud.Barony
}

func (c *多布任DobrzynskaCounty) BDobrzyn多布任() feud.Barony {
	return c.多布任Dobrzyn
}
    
func (c *多布任DobrzynskaCounty) BKikot基库乌() feud.Barony {
	return c.基库乌Kikot
}
    
func (c *多布任DobrzynskaCounty) BLipno利普诺() feud.Barony {
	return c.利普诺Lipno
}
    
func (c *多布任DobrzynskaCounty) BMichalowo米哈沃沃() feud.Barony {
	return c.米哈沃沃Michalowo
}
    
func (c *多布任DobrzynskaCounty) BRypin雷平() feud.Barony {
	return c.雷平Rypin
}
    
func (c *多布任DobrzynskaCounty) BSkepe斯肯佩() feud.Barony {
	return c.斯肯佩Skepe
}
    
func (c *多布任DobrzynskaCounty) BWloclawek奥斯特罗维泰() feud.Barony {
	return c.奥斯特罗维泰Wloclawek
}
    
var CDobrzynska多布任 DobrzynskaCounty = &多布任DobrzynskaCounty{}

func init() {
	f := CDobrzynska多布任.(*多布任DobrzynskaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1581",
		Title:     "dobrzynska",
		TitleName: "多布任",
		TitleCode: "c_dobrzynska",
		Baronies:  map[string]feud.Barony{},
	}

	f.多布任Dobrzyn = BDobrzyn多布任
	f.多布任Dobrzyn.SetParent(f)
	
	f.基库乌Kikot = BKikot基库乌
	f.基库乌Kikot.SetParent(f)
	
	f.利普诺Lipno = BLipno利普诺
	f.利普诺Lipno.SetParent(f)
	
	f.米哈沃沃Michalowo = BMichalowo米哈沃沃
	f.米哈沃沃Michalowo.SetParent(f)
	
	f.雷平Rypin = BRypin雷平
	f.雷平Rypin.SetParent(f)
	
	f.斯肯佩Skepe = BSkepe斯肯佩
	f.斯肯佩Skepe.SetParent(f)
	
	f.奥斯特罗维泰Wloclawek = BWloclawek奥斯特罗维泰
	f.奥斯特罗维泰Wloclawek.SetParent(f)
	
}
