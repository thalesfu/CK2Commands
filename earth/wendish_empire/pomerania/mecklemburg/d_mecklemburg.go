package mecklemburg

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/mecklemburg/hamburg"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/mecklemburg/lubeck"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania/mecklemburg/mecklemburg"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MecklemburgDuke interface {
    feud.Duke
    CHamburg汉堡() 	hamburg.HamburgCounty
    CLubeck吕贝克() 	lubeck.LubeckCounty
    CMecklemburg梅克伦堡() 	mecklemburg.MecklemburgCounty
}

type 梅克伦堡MecklemburgDuke struct {
	feud.BaseDuke
	汉堡Hamburg 	hamburg.HamburgCounty
	吕贝克Lubeck 	lubeck.LubeckCounty
	梅克伦堡Mecklemburg 	mecklemburg.MecklemburgCounty
}

func (d *梅克伦堡MecklemburgDuke) CHamburg汉堡() hamburg.HamburgCounty {
	return d.汉堡Hamburg
}
    
func (d *梅克伦堡MecklemburgDuke) CLubeck吕贝克() lubeck.LubeckCounty {
	return d.吕贝克Lubeck
}
    
func (d *梅克伦堡MecklemburgDuke) CMecklemburg梅克伦堡() mecklemburg.MecklemburgCounty {
	return d.梅克伦堡Mecklemburg
}
    
var DMecklemburg梅克伦堡 MecklemburgDuke = &梅克伦堡MecklemburgDuke{}

func init() {
	f := DMecklemburg梅克伦堡.(*梅克伦堡MecklemburgDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mecklemburg",
		TitleName: "梅克伦堡",
		TitleCode: "d_mecklemburg",
		Counties:  map[string]feud.County{},
	}

	f.汉堡Hamburg = hamburg.CHamburg汉堡
	f.汉堡Hamburg.SetParent(f)
	
	f.吕贝克Lubeck = lubeck.CLubeck吕贝克
	f.吕贝克Lubeck.SetParent(f)
	
	f.梅克伦堡Mecklemburg = mecklemburg.CMecklemburg梅克伦堡
	f.梅克伦堡Mecklemburg.SetParent(f)
	
}
