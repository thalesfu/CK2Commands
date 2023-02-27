package kola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KolaCounty interface {
    feud.County
    BJekanskoi耶坎斯科伊() 	feud.Barony
    BKola科拉() 	feud.Barony
    BLovozero洛沃泽罗() 	feud.Barony
    BMafelskoi马费尔斯科伊() 	feud.Barony
    BMolovskoi莫洛夫斯科伊() 	feud.Barony
    BPechenga佩琴加() 	feud.Barony
    BTre特雷() 	feud.Barony
    BWaranger瓦朗格() 	feud.Barony
}

type 科拉KolaCounty struct {
	feud.BaseCounty
	耶坎斯科伊Jekanskoi 	feud.Barony
	科拉Kola 	feud.Barony
	洛沃泽罗Lovozero 	feud.Barony
	马费尔斯科伊Mafelskoi 	feud.Barony
	莫洛夫斯科伊Molovskoi 	feud.Barony
	佩琴加Pechenga 	feud.Barony
	特雷Tre 	feud.Barony
	瓦朗格Waranger 	feud.Barony
}

func (c *科拉KolaCounty) BJekanskoi耶坎斯科伊() feud.Barony {
	return c.耶坎斯科伊Jekanskoi
}
    
func (c *科拉KolaCounty) BKola科拉() feud.Barony {
	return c.科拉Kola
}
    
func (c *科拉KolaCounty) BLovozero洛沃泽罗() feud.Barony {
	return c.洛沃泽罗Lovozero
}
    
func (c *科拉KolaCounty) BMafelskoi马费尔斯科伊() feud.Barony {
	return c.马费尔斯科伊Mafelskoi
}
    
func (c *科拉KolaCounty) BMolovskoi莫洛夫斯科伊() feud.Barony {
	return c.莫洛夫斯科伊Molovskoi
}
    
func (c *科拉KolaCounty) BPechenga佩琴加() feud.Barony {
	return c.佩琴加Pechenga
}
    
func (c *科拉KolaCounty) BTre特雷() feud.Barony {
	return c.特雷Tre
}
    
func (c *科拉KolaCounty) BWaranger瓦朗格() feud.Barony {
	return c.瓦朗格Waranger
}
    
var CKola科拉 KolaCounty = &科拉KolaCounty{}

func init() {
	f := CKola科拉.(*科拉KolaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "387",
		Title:     "kola",
		TitleName: "科拉",
		TitleCode: "c_kola",
		Baronies:  map[string]feud.Barony{},
	}

	f.耶坎斯科伊Jekanskoi = BJekanskoi耶坎斯科伊
	f.耶坎斯科伊Jekanskoi.SetParent(f)
	
	f.科拉Kola = BKola科拉
	f.科拉Kola.SetParent(f)
	
	f.洛沃泽罗Lovozero = BLovozero洛沃泽罗
	f.洛沃泽罗Lovozero.SetParent(f)
	
	f.马费尔斯科伊Mafelskoi = BMafelskoi马费尔斯科伊
	f.马费尔斯科伊Mafelskoi.SetParent(f)
	
	f.莫洛夫斯科伊Molovskoi = BMolovskoi莫洛夫斯科伊
	f.莫洛夫斯科伊Molovskoi.SetParent(f)
	
	f.佩琴加Pechenga = BPechenga佩琴加
	f.佩琴加Pechenga.SetParent(f)
	
	f.特雷Tre = BTre特雷
	f.特雷Tre.SetParent(f)
	
	f.瓦朗格Waranger = BWaranger瓦朗格
	f.瓦朗格Waranger.SetParent(f)
	
}
