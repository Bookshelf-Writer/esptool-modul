package esp32

import (
	"github.com/Bookshelf-Writer/esptool-modul/esp32/code"
	"strings"
)

type Features map[code.FeatureType]bool

func (f Features) String() string {
	res := []string{}
	for feature, status := range f {
		if !status {
			continue
		}
		res = append(res, feature.String())
	}
	return strings.Join(res, ", ")
}

func (e *ESP32ROM) GetFeatures() (Features, error) {
	features := Features{
		code.FeatureWiFi: true,
	}

	word3, err := ReadEfuse(e.SerialPort, e.defaultTimeout, 3)
	if err != nil {
		return features, err
	}

	features[code.FeatureBluetooth] = word3[0]&(1<<1) == 0
	features[code.FeatureDualCore] = word3[0]&(1<<0) > 0
	features[code.FeatureSingleCore] = !features[code.FeatureDualCore]
	if word3[1]&(1<<5) > 0 {
		features[code.FeatureClock160MHz] = word3[1]&(1<<4) > 0
		features[code.FeatureClock240MHz] = !features[code.FeatureClock160MHz]
	}

	pkgVersion := (word3[1] >> 1) & 0x07
	features[code.FeatureEmbeddedFlash] = pkgVersion == 2 || pkgVersion == 4 || pkgVersion == 5

	word4, err := ReadEfuse(e.SerialPort, e.defaultTimeout, 4)
	if err != nil {
		return features, err
	}

	features[code.FeatureVRefCalibration] = word4[1]&0x1F > 0
	features[code.FeatureBLK3PartiallyReserved] = word4[1]>>6&0x01 > 0

	word6, err := ReadEfuse(e.SerialPort, e.defaultTimeout, 6)
	if err != nil {
		return features, err
	}

	features[code.FeatureCodingSchemeNone] = word6[0]&0x03 == 0
	features[code.FeatureCodingScheme34] = word6[0]&0x03 == 1
	features[code.FeatureCodingSchemeRepeat] = word6[0]&0x03 == 2
	features[code.FeatureCodingSchemeInvalid] = word6[0]&0x03 == 3

	return features, nil
}
