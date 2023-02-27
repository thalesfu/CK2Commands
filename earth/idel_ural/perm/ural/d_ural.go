package ural

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/ural/thisageta"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/ural/ural"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/ural/vishera"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UralDuke interface {
    feud.Duke
    CThisageta杜撒该塔伊() 	thisageta.ThisagetaCounty
    CUral乌拉尔() 	ural.UralCounty
    CVishera维舍拉() 	vishera.VisheraCounty
}

type 乌拉尔UralDuke struct {
	feud.BaseDuke
	杜撒该塔伊Thisageta 	thisageta.ThisagetaCounty
	乌拉尔Ural 	ural.UralCounty
	维舍拉Vishera 	vishera.VisheraCounty
}

func (d *乌拉尔UralDuke) CThisageta杜撒该塔伊() thisageta.ThisagetaCounty {
	return d.杜撒该塔伊Thisageta
}
    
func (d *乌拉尔UralDuke) CUral乌拉尔() ural.UralCounty {
	return d.乌拉尔Ural
}
    
func (d *乌拉尔UralDuke) CVishera维舍拉() vishera.VisheraCounty {
	return d.维舍拉Vishera
}
    
var DUral乌拉尔 UralDuke = &乌拉尔UralDuke{}

func init() {
	f := DUral乌拉尔.(*乌拉尔UralDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ural",
		TitleName: "乌拉尔",
		TitleCode: "d_ural",
		Counties:  map[string]feud.County{},
	}

	f.杜撒该塔伊Thisageta = thisageta.CThisageta杜撒该塔伊
	f.杜撒该塔伊Thisageta.SetParent(f)
	
	f.乌拉尔Ural = ural.CUral乌拉尔
	f.乌拉尔Ural.SetParent(f)
	
	f.维舍拉Vishera = vishera.CVishera维舍拉
	f.维舍拉Vishera.SetParent(f)
	
}
