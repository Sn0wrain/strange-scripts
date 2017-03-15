package main

import "fmt"

const (
	// Get location by gps cords
	locationApi = `http://api.map.baidu.com/geocoder/v2/?callback=renderReverse&location=%s,%s&output=json&pois=1&ak=WunXh7fwhTk3Z1OcnL89ylvdmiZcVS0L`
	// conver cord (lon lat)
	cordConvertApi = `http://api.map.baidu.com/geoconv/v1/?coords=%s,%s&ak=WunXh7fwhTk3Z1OcnL89ylvdmiZcVS0L&from=%s&output=json`
)

func main() {
}

func gps2Baidu(lngGps, latGps string) (string, string, error) {
	lng := ""
	lat := ""
	resp, err := http.Get(fmt.Sprintf(
		cordConvertApi,
		lngGps,
		latGps,
		"1",
	))
	if err != nil {
		log.Warn(errors.As(err))
		return lng, lat, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn(errors.As(err))
		return lng, lat, err
	}
	bodyStr := B2S(body)
	bodyStrs := strings.SplitAfterN(bodyStr, ": ", 2)
	jStr := bodyStrs[0]
	jByte := S2B(jStr)

	cResp := convResp{}
	if err := json.Unmarshal(jByte, &cResp); err != nil {
		log.Warn(errors.As(err))
		return lng, lat, err
	}

	for _, v := range cResp.Result {
		lng = strconv.FormatFloat(v.X, 'f', -1, 64)
		lat = strconv.FormatFloat(v.Y, 'f', -1, 64)
	}
	return lng, lat, nil
}

// 获取位置从HTML换为腾讯SDK
func ten2Baidu(lngTen, latTen string) (string, string, error) {
	lng := ""
	lat := ""
	resp, err := http.Get(fmt.Sprintf(
		cordConvertApi,
		lngTen,
		latTen,
		"3",
	))
	if err != nil {
		log.Warn(errors.As(err))
		return lng, lat, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn(errors.As(err))
		return lng, lat, err
	}
	bodyStr := B2S(body)
	bodyStrs := strings.SplitAfterN(bodyStr, ": ", 2)
	jStr := bodyStrs[0]
	jByte := S2B(jStr)

	cResp := convResp{}
	if err := json.Unmarshal(jByte, &cResp); err != nil {
		log.Warn(errors.As(err))
		return lng, lat, err
	}

	for _, v := range cResp.Result {
		lng = strconv.FormatFloat(v.X, 'f', -1, 64)
		lat = strconv.FormatFloat(v.Y, 'f', -1, 64)
	}
	return lng, lat, nil
}

func getCityByLocation(lng, lat string) (*model.City, error) {
	resp, err := http.Get(fmt.Sprintf(
		locationApi,
		lat,
		lng,
	))
	if err != nil {
		log.Warn(errors.As(err))
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn(errors.As(err))
		return nil, err
	}
	bodyStr := B2S(body)
	bodyStrs := strings.SplitAfterN(bodyStr, "(", 2)
	bodyB := S2B(bodyStrs[1])
	jStr := B2S(bodyB[0 : len(bodyB)-1])
	jByte := S2B(jStr)

	mResp := MapResp{}
	if err := json.Unmarshal(jByte, &mResp); err != nil {
		log.Warn(errors.As(err))
		return nil, err
	}
	lifeDB := model.NewLifeDB()
	return lifeDB.GetCityByName(mResp.Result.AddressComponent.City)
}

// DEPRECATED
func getDistance2(lngX, latX, lngY, latY float64) string {
	c := math.Sin(latX)*math.Sin(latY)*math.Cos(lngX-lngY) + math.Cos(latX)*math.Cos(latY)
	dis := 6371.004 * math.Acos(c) * math.Pi / 180
	if dis < 1 {
		return strconv.Itoa((int)(dis*1000)) + "m"
	} else {
		fmt.Println("dis:", dis)
		return fmt.Sprintf("%.1f", dis) + "km"
	}
}

func getDistance(lngX, latX, lngY, latY float64) string {
	lat1 := (math.Pi / 180) * latX
	lat2 := (math.Pi / 180) * latY
	lng1 := (math.Pi / 180) * lngX
	lng2 := (math.Pi / 180) * lngY
	R := 6371.004
	dis := math.Acos(math.Sin(lat1)*math.Sin(lat2)+math.Cos(lat1)*math.Cos(lat2)*math.Cos(lng2-lng1)) * R
	if dis < 1 {
		return strconv.Itoa((int)(dis*1000)) + "m"
	} else {
		return fmt.Sprintf("%.1f", dis) + "km"
	}
}
