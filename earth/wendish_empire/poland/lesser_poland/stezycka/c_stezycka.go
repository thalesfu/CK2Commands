package stezycka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type StezyckaCounty interface {
    feud.County
    BDeblin登布林() 	feud.Barony
    BKock科茨克() 	feud.Barony
    BLukow武库夫() 	feud.Barony
    BRadzyn拉曾() 	feud.Barony
    BRyki雷基() 	feud.Barony
    BStezyca斯滕日察() 	feud.Barony
    BStoczeklukowski武库夫地区斯托切克() 	feud.Barony
}

type 斯滕日察StezyckaCounty struct {
	feud.BaseCounty
	登布林Deblin 	feud.Barony
	科茨克Kock 	feud.Barony
	武库夫Lukow 	feud.Barony
	拉曾Radzyn 	feud.Barony
	雷基Ryki 	feud.Barony
	斯滕日察Stezyca 	feud.Barony
	武库夫地区斯托切克Stoczeklukowski 	feud.Barony
}

func (c *斯滕日察StezyckaCounty) BDeblin登布林() feud.Barony {
	return c.登布林Deblin
}
    
func (c *斯滕日察StezyckaCounty) BKock科茨克() feud.Barony {
	return c.科茨克Kock
}
    
func (c *斯滕日察StezyckaCounty) BLukow武库夫() feud.Barony {
	return c.武库夫Lukow
}
    
func (c *斯滕日察StezyckaCounty) BRadzyn拉曾() feud.Barony {
	return c.拉曾Radzyn
}
    
func (c *斯滕日察StezyckaCounty) BRyki雷基() feud.Barony {
	return c.雷基Ryki
}
    
func (c *斯滕日察StezyckaCounty) BStezyca斯滕日察() feud.Barony {
	return c.斯滕日察Stezyca
}
    
func (c *斯滕日察StezyckaCounty) BStoczeklukowski武库夫地区斯托切克() feud.Barony {
	return c.武库夫地区斯托切克Stoczeklukowski
}
    
var CStezycka斯滕日察 StezyckaCounty = &斯滕日察StezyckaCounty{}

func init() {
	f := CStezycka斯滕日察.(*斯滕日察StezyckaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1584",
		Title:     "stezycka",
		TitleName: "斯滕日察",
		TitleCode: "c_stezycka",
		Baronies:  map[string]feud.Barony{},
	}

	f.登布林Deblin = BDeblin登布林
	f.登布林Deblin.SetParent(f)
	
	f.科茨克Kock = BKock科茨克
	f.科茨克Kock.SetParent(f)
	
	f.武库夫Lukow = BLukow武库夫
	f.武库夫Lukow.SetParent(f)
	
	f.拉曾Radzyn = BRadzyn拉曾
	f.拉曾Radzyn.SetParent(f)
	
	f.雷基Ryki = BRyki雷基
	f.雷基Ryki.SetParent(f)
	
	f.斯滕日察Stezyca = BStezyca斯滕日察
	f.斯滕日察Stezyca.SetParent(f)
	
	f.武库夫地区斯托切克Stoczeklukowski = BStoczeklukowski武库夫地区斯托切克
	f.武库夫地区斯托切克Stoczeklukowski.SetParent(f)
	
}
