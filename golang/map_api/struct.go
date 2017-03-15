package main

import (
	"reflect"
	"unsafe"
)

type convResp struct {
	Status int       `json:"status"`
	Result []ConvRes `json:"result"`
}

type ConvRes struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type MapResp struct {
	Status int `json:"status"`
	Result Res `json:"result"`
}

type Res struct {
	Location           Loc    `json:"location"`
	FormattedAddress   string `json:"formatted_address"`
	Business           string `json:"business"`
	AddressComponent   AComp  `json:"addressComponent"`
	Pois               []Poi  `json:"pois"`
	PoiRegions         []PoiR `json:"poiRegions"`
	SematicDescription string `json:"sematic_description"`
	CityCode           int    `json:"cityCode"`
}

type Loc struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

type AComp struct {
	Country     string `json:"country"`
	CountryCode int    `json:"country_code"`
	Province    string `json:"province"`
	City        string `json:"city"`
	District    string `json:"district"`
	Adcode      string `json:""adcode"`
	Street      string `json:"street"`
	StreetNum   string `json:"street_number"`
	Direction   string `json:"direction"`
	Distance    string `json:"distance"`
}

type Poi struct {
	Addr      string `json:"addr"`
	Cp        string `json:"cp"`
	Direction string `json:"direction"`
	Distance  string `json:"distance"`
	Name      string `json:"name"`
	PoiType   string `json:"poiType"`
	Pnt       Point  `json:"point"`
	Tag       string `json:"tag"`
	Tel       string `json:"tel"`
	Uid       string `json:"uid"`
	Zip       string `json:"zip"`
}

type PoiR struct {
	DirectionDesc string `json:"direction_desc"`
	Name          string `json:"name"`
	Tag           string `json:"tag"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func B2S(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

func S2B(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}
