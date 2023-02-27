package samos

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/samos/chios"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/samos/ephesos"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/samos/smyrna"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SamosDuke interface {
    feud.Duke
    CChios希俄斯() 	chios.ChiosCounty
    CEphesos以弗所() 	ephesos.EphesosCounty
    CSmyrna士麦拿() 	smyrna.SmyrnaCounty
}

type 萨摩斯SamosDuke struct {
	feud.BaseDuke
	希俄斯Chios 	chios.ChiosCounty
	以弗所Ephesos 	ephesos.EphesosCounty
	士麦拿Smyrna 	smyrna.SmyrnaCounty
}

func (d *萨摩斯SamosDuke) CChios希俄斯() chios.ChiosCounty {
	return d.希俄斯Chios
}
    
func (d *萨摩斯SamosDuke) CEphesos以弗所() ephesos.EphesosCounty {
	return d.以弗所Ephesos
}
    
func (d *萨摩斯SamosDuke) CSmyrna士麦拿() smyrna.SmyrnaCounty {
	return d.士麦拿Smyrna
}
    
var DSamos萨摩斯 SamosDuke = &萨摩斯SamosDuke{}

func init() {
	f := DSamos萨摩斯.(*萨摩斯SamosDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "samos",
		TitleName: "萨摩斯",
		TitleCode: "d_samos",
		Counties:  map[string]feud.County{},
	}

	f.希俄斯Chios = chios.CChios希俄斯
	f.希俄斯Chios.SetParent(f)
	
	f.以弗所Ephesos = ephesos.CEphesos以弗所
	f.以弗所Ephesos.SetParent(f)
	
	f.士麦拿Smyrna = smyrna.CSmyrna士麦拿
	f.士麦拿Smyrna.SetParent(f)
	
}
