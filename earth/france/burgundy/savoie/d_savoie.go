package savoie

import (
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/savoie/aosta"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/savoie/geneve"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/savoie/savoie"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/savoie/valais"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SavoieDuke interface {
    feud.Duke
    CAosta奥斯塔() 	aosta.AostaCounty
    CGeneve日内瓦() 	geneve.GeneveCounty
    CSavoie萨伏依() 	savoie.SavoieCounty
    CValais瓦莱() 	valais.ValaisCounty
}

type 萨伏依SavoieDuke struct {
	feud.BaseDuke
	奥斯塔Aosta 	aosta.AostaCounty
	日内瓦Geneve 	geneve.GeneveCounty
	萨伏依Savoie 	savoie.SavoieCounty
	瓦莱Valais 	valais.ValaisCounty
}

func (d *萨伏依SavoieDuke) CAosta奥斯塔() aosta.AostaCounty {
	return d.奥斯塔Aosta
}
    
func (d *萨伏依SavoieDuke) CGeneve日内瓦() geneve.GeneveCounty {
	return d.日内瓦Geneve
}
    
func (d *萨伏依SavoieDuke) CSavoie萨伏依() savoie.SavoieCounty {
	return d.萨伏依Savoie
}
    
func (d *萨伏依SavoieDuke) CValais瓦莱() valais.ValaisCounty {
	return d.瓦莱Valais
}
    
var DSavoie萨伏依 SavoieDuke = &萨伏依SavoieDuke{}

func init() {
	f := DSavoie萨伏依.(*萨伏依SavoieDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "savoie",
		TitleName: "萨伏依",
		TitleCode: "d_savoie",
		Counties:  map[string]feud.County{},
	}

	f.奥斯塔Aosta = aosta.CAosta奥斯塔
	f.奥斯塔Aosta.SetParent(f)
	
	f.日内瓦Geneve = geneve.CGeneve日内瓦
	f.日内瓦Geneve.SetParent(f)
	
	f.萨伏依Savoie = savoie.CSavoie萨伏依
	f.萨伏依Savoie.SetParent(f)
	
	f.瓦莱Valais = valais.CValais瓦莱
	f.瓦莱Valais.SetParent(f)
	
}
