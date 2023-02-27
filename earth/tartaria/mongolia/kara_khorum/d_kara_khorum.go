package kara_khorum

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/kara_khorum/kara_khorum"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/kara_khorum/orkhon"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/kara_khorum/tuul"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Kara_khorumDuke interface {
    feud.Duke
    CKara_khorum哈剌和林() 	kara_khorum.Kara_khorumCounty
    COrkhon嗢昆() 	orkhon.OrkhonCounty
    CTuul独乐() 	tuul.TuulCounty
}

type 哈剌和林Kara_khorumDuke struct {
	feud.BaseDuke
	哈剌和林Kara_khorum 	kara_khorum.Kara_khorumCounty
	嗢昆Orkhon 	orkhon.OrkhonCounty
	独乐Tuul 	tuul.TuulCounty
}

func (d *哈剌和林Kara_khorumDuke) CKara_khorum哈剌和林() kara_khorum.Kara_khorumCounty {
	return d.哈剌和林Kara_khorum
}
    
func (d *哈剌和林Kara_khorumDuke) COrkhon嗢昆() orkhon.OrkhonCounty {
	return d.嗢昆Orkhon
}
    
func (d *哈剌和林Kara_khorumDuke) CTuul独乐() tuul.TuulCounty {
	return d.独乐Tuul
}
    
var DKara_khorum哈剌和林 Kara_khorumDuke = &哈剌和林Kara_khorumDuke{}

func init() {
	f := DKara_khorum哈剌和林.(*哈剌和林Kara_khorumDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kara_khorum",
		TitleName: "哈剌和林",
		TitleCode: "d_kara_khorum",
		Counties:  map[string]feud.County{},
	}

	f.哈剌和林Kara_khorum = kara_khorum.CKara_khorum哈剌和林
	f.哈剌和林Kara_khorum.SetParent(f)
	
	f.嗢昆Orkhon = orkhon.COrkhon嗢昆
	f.嗢昆Orkhon.SetParent(f)
	
	f.独乐Tuul = tuul.CTuul独乐
	f.独乐Tuul.SetParent(f)
	
}
