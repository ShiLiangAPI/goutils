package fileutils

import "github.com/ShiLiangAPI/goutils/function"

func GetContentTypeBySuffix(suffix string) string {
	imgList := []string{"jpeg", "jpg", "png", "gif", "tif", "bmp", "dwg"}
	exists := function.ValueInSlice[string](imgList, suffix)
	if exists {
		return "IMAGE"
	}

	audioList := []string{"mp3", "wma", "wav", "mid", "ape", "flac"}
	existAudio := function.ValueInSlice[string](audioList, suffix)
	if existAudio {
		return "AUDIO"
	}

	videoList := []string{"rmvb", "flv", "mp4", "mpg", "mpeg", "avi", "rm", "mov", "wmv", "webm"}
	existVideo := function.ValueInSlice[string](videoList, suffix)
	if existVideo {
		return "VIDEO"
	}

	return "FILE"
}
