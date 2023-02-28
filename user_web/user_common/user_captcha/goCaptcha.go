package user_captcha

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wenlng/go-captcha/captcha"
	"golang.org/x/image/font"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type GoCaptcha struct {
	Dots        map[int]captcha.CharDot `json:"dots"`
	Base64      string                  `json:"base64"`
	ThumbBase64 string                  `json:"thumb_base64"`
	Key         string                  `json:"key"`
}

type GoCaptchaResponse struct {
	Base64      string `json:"base64" binding:"required"`
	ThumbBase64 string `json:"thumb_base64" binding:"required"`
	Key         string `json:"key" binding:"required"`
}

type CheckCaptchaRequest struct {
	Dots string `json:"dots"`
	Key  string `json:"key"`
}

func GenBigCaptcha() (*GoCaptcha, error) {
	capt := captcha.GetCaptcha()
	path, _ := os.Getwd()
	println("path:", path)
	capt.SetFont([]string{
		path + "/resources/fonts/fzshengsksjw_cu.ttf",
	})
	capt.SetBackground([]string{
		path + "/resources/images/1.jpg",
		path + "/resources/images/2.jpg",
	})
	capt.SetImageSize(captcha.Size{300, 300})
	capt.SetImageQuality(captcha.QualityCompressNone)
	capt.SetFontHinting(font.HintingFull)
	capt.SetTextRangLen(captcha.RangeVal{6, 7})
	capt.SetRangFontSize(captcha.RangeVal{32, 42})
	capt.SetTextRangFontColors([]string{
		"#1d3f84",
		"#3a6a1e",
	})
	capt.SetImageFontAlpha(0.5)
	capt.SetTextShadow(true)
	capt.SetTextShadowColor("#101010")
	capt.SetTextShadowPoint(captcha.Point{1, 1})
	capt.SetTextRangAnglePos([]captcha.RangeVal{
		{1, 15},
		{15, 30},
		{30, 45},
		{315, 330},
		{330, 345},
		{345, 359},
	})
	capt.SetImageFontDistort(captcha.DistortLevel2)
	//chars := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//_ = capt.SetRangChars(strings.Split(chars, ""))
	//chars := []string{"HE","CA","WO","NE","HT","IE","PG","GI","CH","CO","DA"}
	//_ = capt.SetRangChars(chars)
	chars := []string{"你", "好", "呀", "这", "是", "点", "击", "验", "证", "码", "哟"}
	_ = capt.SetRangChars(chars)
	dots, imageBase64, thumbImageBase64, key, err := capt.Generate()
	if err != nil {
		return nil, err
	}
	fmt.Println(len(imageBase64))
	fmt.Println(len(thumbImageBase64))
	fmt.Println(key)
	fmt.Println(dots)
	goCaptcha := &GoCaptcha{
		Dots:        dots,
		Base64:      imageBase64,
		ThumbBase64: thumbImageBase64,
		Key:         key,
	}
	return goCaptcha, nil
}

func GenSmallCaptcha() (*GoCaptcha, error) {
	capt := captcha.GetCaptcha()

	path, _ := os.Getwd()

	// ====================================================
	// Method: SetThumbSize(size Size);
	// Desc: 设置缩略图的尺寸
	// ====================================================
	capt.SetThumbSize(captcha.Size{150, 40})

	// ====================================================
	// Method: SetRangCheckTextLen(val captcha.RangeVal);
	// Desc:设置缩略图校验文本的随机长度范围
	// ====================================================
	capt.SetRangCheckTextLen(captcha.RangeVal{2, 4})

	// ====================================================
	// Method: SetRangCheckFontSize(val captcha.RangeVal);
	// Desc:设置缩略图校验文本的随机大小
	// ====================================================
	capt.SetRangCheckFontSize(captcha.RangeVal{24, 30})

	// ====================================================
	// Method: SetThumbTextRangFontColors(colors []string);
	// Desc: 设置缩略图文本的随机十六进制颜色
	// ====================================================
	capt.SetThumbTextRangFontColors([]string{
		"#1d3f84",
		"#3a6a1e",
	})

	// ====================================================
	// Method: SetThumbBgColors(colors []string);
	// Desc: 设置缩略图的背景随机十六进制颜色
	// ====================================================
	capt.SetThumbBgColors([]string{
		"#1d3f84",
		"#3a6a1e",
	})

	// ====================================================
	// Method: SetThumbBackground(colors []string);
	// Desc:设置缩略图的随机图像背景，自动仅读取一次并加载到内存中缓存，如需重置可清除缓存
	// ====================================================
	capt.SetThumbBackground([]string{
		path + "/resources/images/r1.jpg",
		path + "/resources/images/r2.jpg",
	})

	// ====================================================
	// Method: SetThumbBgDistort(val int);
	// Desc:设置缩略图背景的扭曲程度
	// ====================================================
	capt.SetThumbBgDistort(captcha.DistortLevel2)

	// ====================================================
	// Method: SetThumbFontDistort(val int);
	// Desc:设置缩略图字体的扭曲程度
	// ====================================================
	capt.SetThumbFontDistort(captcha.DistortLevel2)

	// ====================================================
	// Method: SetThumbBgCirclesNum(val int);
	// Desc:设置缩略图背景的圈点数
	// ====================================================
	capt.SetThumbBgCirclesNum(20)

	// ====================================================
	// Method: SetThumbBgSlimLineNum(val int);
	// Desc:设置缩略图背景的线条数
	// ====================================================
	capt.SetThumbBgSlimLineNum(3)
	chars := []string{"你", "好", "呀", "这", "是", "点", "击", "验", "证", "码", "哟"}
	_ = capt.SetRangChars(chars)
	dots, imageBase64, thumbImageBase64, key, err := capt.Generate()
	if err != nil {
		return nil, err
	}
	fmt.Println(len(imageBase64))
	fmt.Println(len(thumbImageBase64))
	fmt.Println(key)
	fmt.Println(dots)
	goCaptcha := &GoCaptcha{
		Dots:        dots,
		Base64:      imageBase64,
		ThumbBase64: thumbImageBase64,
		Key:         key,
	}
	writeCache(dots, key)
	return goCaptcha, nil
}

func VerifyCaptcha(request *CheckCaptchaRequest) (checked bool, err error) {
	dots := request.Dots
	key := request.Key
	checked = false
	if dots == "" || key == "" {
		err = errors.New("dots or key param is empty")
		return
	}
	cacheData := readCache(key)
	if cacheData == "" {
		err = errors.New("illegal key")
		return
	}
	src := strings.Split(dots, ",")

	var dct map[int]captcha.CharDot
	if err = json.Unmarshal([]byte(cacheData), &dct); err != nil {
		err = errors.New("illegal key")
		return
	}
	chkRet := false
	if (len(dct) * 2) == len(src) {
		for i, dot := range dct {
			j := i * 2
			k := i*2 + 1
			sx, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[j]), 64)
			sy, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[k]), 64)

			// 检测点位置
			// chkRet = captcha.CheckPointDist(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height))

			// 校验点的位置,在原有的区域上添加额外边距进行扩张计算区域,不推荐设置过大的padding
			// 例如：文本的宽和高为30，校验范围x为10-40，y为15-45，此时扩充5像素后校验范围宽和高为40，则校验范围x为5-45，位置y为10-50
			chkRet = captcha.CheckPointDistWithPadding(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height), 5)
			if !chkRet {
				break
			}
		}
	}
	if chkRet {
		// 通过校验
		checked = true
	}
	return
}

/**
 * @Description: Write Cache，“Redis” is recommended
 * @param v
 * @param file
 */
func writeCache(v interface{}, file string) {
	bt, _ := json.Marshal(v)
	month := time.Now().Month().String()
	cacheDir := getCacheDir() + month + "/"
	_ = os.MkdirAll(cacheDir, 0660)
	file = cacheDir + file + ".json"
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer logFile.Close()
	// 检查过期文件
	//checkCacheOvertimeFile()
	_, _ = io.WriteString(logFile, string(bt))
}

/**
 * @Description: Read Cache，“Redis” is recommended
 * @param file
 * @return string
 */
func readCache(file string) string {
	month := time.Now().Month().String()
	cacheDir := getCacheDir() + month + "/"
	file = cacheDir + file + ".json"

	if !checkFileIsExist(file) {
		return ""
	}

	bt, err := os.ReadFile(file)
	err = os.Remove(file)
	if err == nil {
		return string(bt)
	}
	return ""
}

/**
 * @Description: Calculate the distance between two points
 * @param sx
 * @param sy
 * @param dx
 * @param dy
 * @param width
 * @param height
 * @return bool
 */
func checkDist(sx, sy, dx, dy, width, height int64) bool {
	return sx >= dx &&
		sx <= dx+width &&
		sy <= dy &&
		sy >= dy-height
}

/**
 * @Description: Get cache dir path
 * @return string
 */
func getCacheDir() string {
	return getPWD() + "/.cache/"
}

/**
 * @Description: Get pwd dir path
 * @return string
 */
func getPWD() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

/**
 * @Description: Check file exist
 * @param filename
 * @return bool
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/**
 * @Description: 启动定时任务, 5分钟执行一次
 */
func runTimedTask() {
	ticker := time.NewTicker(time.Minute * 5)
	go func() {
		for range ticker.C {
			checkCacheOvertimeFile()
		}
	}()
}

func getFileCreateTime(path string) int64 {
	// osType := runtime.GOOS
	fileInfo, _ := os.Stat(path)
	statT := fileInfo.Sys().(*syscall.Stat_t)
	tCreate, _ := statT.Ctimespec.Unix()
	return tCreate
}

/**
 * @Description: 检查缓存超时文件， 30分钟
 */
func checkCacheOvertimeFile() {
	files, files1, _ := listDir(getCacheDir())
	for _, table := range files1 {
		temp, _, _ := listDir(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	for _, file := range files {
		t := getFileCreateTime(file)
		ex := time.Now().Unix() - t
		if ex > (60 * 30) {
			_ = os.Remove(file)
		}
	}
}

/**
 * @Description: 获取目录文件列表
 * @param dirPth
 * @return files
 * @return files1
 * @return err
 */
func listDir(dirPth string) (files []string, files1 []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			files1 = append(files1, dirPth+PthSep+fi.Name())
			_, _, _ = listDir(dirPth + PthSep + fi.Name())
		} else {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, files1, nil
}
