package ua

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-yaml/yaml"
)

// UserAgentParser 查询浏览器的类结构
type UserAgentParser struct {
	UserAgentRe       string         `yaml:"regex"`
	FamilyReplaceMent string         `yaml:"family_replacement,omitempty"`
	V1ReplaceMent     string         `yaml:"v1_replacement,omitempty"`
	V2ReplaceMent     string         `yaml:"v2_replacement,omitempty"`
	Reg               *regexp.Regexp `yaml:"my ,omitempty"`
}

// OSParser 查询系统的类结构
type OSParser struct {
	UserAgentRe     string         `yaml:"regex"`
	OSReplaceMent   string         `yaml:"os_replacement,omitempty"`
	OSV1ReplaceMent string         `yaml:"os_v1_replacement,omitempty"`
	OSV2ReplaceMent string         `yaml:"os_v2_replacement,omitempty"`
	Reg             *regexp.Regexp `yaml:"my ,omitempty"`
}

// DeviceParser 查询设备的类结构
type DeviceParser struct {
	UserAgentRe       string         `yaml:"regex"`
	RegexFlag         string         `yaml:"regex_flag,omitempty"`
	DeviceReplaceMent string         `yaml:"device_replacement,omitempty"`
	BrandReplaceMent  string         `yaml:"brand_replacement,omitempty"`
	ModelReplaceMent  string         `yaml:"model_replacement,omitempty"`
	Reg               *regexp.Regexp `yaml:"my ,omitempty"`
}

// DeviceResult 返回的设备信息
type DeviceResult struct {
	Family string
	Brand  string
	Model  string
}

// UAResult 返回的浏览器信息
type UAResult struct {
	Family string
	Major  string
	Minor  string
	Patch  string
}

// OSResult 返回的系统信息
type OSResult struct {
	Family     string
	Major      string
	Minor      string
	Patch      string
	PatchMinor string
}

// T 解析结构
type T struct {
	UserAgentParsers []*UserAgentParser `yaml:"user_agent_parsers" `
	OsParsers        []*OSParser        `yaml:"os_parsers"`
	DeviceParsers    []*DeviceParser    `yaml:"device_parsers"`
}

//Create USER_AGENT_PARSERS, OS_PARSERS, and DEVICE_PARSERS arrays
func loadyml() ([]*DeviceParser, []*OSParser, []*UserAgentParser) {

	t := T{}

	// name, err := os.Open("/Users/mac/go/src/uaparse/uaparses/regexes.yaml")
	// pri(err)
	//
	// data, err := ioutil.ReadAll(name)
	// pri(err)

	err := yaml.Unmarshal([]byte(Data), &t)
	pri(err)

	for _, i := range t.DeviceParsers {
		i.Reg = regexp.MustCompile(i.UserAgentRe)
	}

	for _, i := range t.OsParsers {
		i.Reg = regexp.MustCompile(i.UserAgentRe)
	}

	for _, i := range t.UserAgentParsers {
		i.Reg = regexp.MustCompile(i.UserAgentRe)
	}

	return t.DeviceParsers, t.OsParsers, t.UserAgentParsers
}

// const
var (
	deviceParsers, osParsers, uaParsers = loadyml()
	renum                               = regexp.MustCompile("\\$(\\d)")
	respace                             = regexp.MustCompile("^\\s+|\\s+\\$")
)

// innerReplace helper function for parsedevice
func innerReplace(str string, group []string) string {
	idx, err := strconv.Atoi(str[1:])
	if err != nil {
		panic(err)
	}
	idx--
	if idx < len(group) {
		return group[idx]
	}
	return ""
}

// multiReplace helper function for parsedevice
func multiReplace(str string, mtch []string) string {
	res := renum.FindAllString(str, -1)
	if len(res) == 0 {
		return str
	}
	for _, i := range res {
		str = strings.ReplaceAll(str, i, innerReplace(i, mtch))
	}
	return respace.ReplaceAllString(str, "")
}

// ParseDevice 设备解析函数
func ParseDevice(userAgent string) *DeviceResult {

	device, brand, model := "", "", ""
	for _, i := range deviceParsers {

		strs := i.Reg.FindStringSubmatch(userAgent)

		if len(strs) > 1 {

			strs = strs[1:]

			if i.DeviceReplaceMent != "" {
				device = multiReplace(i.DeviceReplaceMent, strs)
			} else {
				device = strs[0]
			}

			if i.BrandReplaceMent != "" {
				brand = multiReplace(i.BrandReplaceMent, strs)
			} else {
				brand = strs[1]
			}

			if i.ModelReplaceMent != "" {
				model = multiReplace(i.ModelReplaceMent, strs)
			} else if len(strs) > 1 {
				model = strs[2]
			}

			break

		}

	}
	return &DeviceResult{device, brand, model}
}

// ParseUA UA解析函数
func ParseUA(userAgent string) *UAResult {
	family, major, minor, patch := "", "", "", ""
	for _, value := range uaParsers {

		strs := value.Reg.FindStringSubmatch(userAgent)

		if len(strs) > 1 {

			strs = strs[1:]
			if value.FamilyReplaceMent != "" {
				if ok, _ := regexp.Match("$1", []byte(value.FamilyReplaceMent)); ok {
					family = strings.ReplaceAll(value.FamilyReplaceMent, "$1", strs[0])
				} else {
					family = value.FamilyReplaceMent
				}
			} else {
				family = strs[0]
			}
			if value.V1ReplaceMent != "" {
				major = value.V1ReplaceMent
			} else if len(strs) > 1 {
				major = strs[1]
			}
			if value.V2ReplaceMent != "" {
				minor = value.V2ReplaceMent
			} else if len(strs) > 2 {
				minor = strs[2]
			}
			if len(strs) > 3 {
				patch = strs[3]
			}
		}
	}
	return &UAResult{family, major, minor, patch}
}

// ParseOS OS解析函数
func ParseOS(userAgent string) *OSResult {

	family, major, minor, patch, patchminor := "", "", "", "", ""
	for _, value := range osParsers {

		strs := value.Reg.FindStringSubmatch(userAgent)

		if len(strs) > 1 {

			strs = strs[1:]
			if value.OSReplaceMent != "" {
				family = value.OSReplaceMent
			} else {
				family = strs[0]
			}
			if value.OSV1ReplaceMent != "" {
				major = value.OSV1ReplaceMent
			} else if len(strs) > 1 {
				major = strs[1]
			}
			if value.OSV2ReplaceMent != "" {
				minor = value.OSV2ReplaceMent
			} else if len(strs) > 2 {
				minor = strs[2]
			}
			if len(strs) > 3 {
				patch = strs[3]
			}
			if len(strs) > 4 {
				patchminor = strs[4]
			}
			break
		}

	}

	return &OSResult{family, major, minor, patch, patchminor}
}

func pri(err error) {
	if err != nil {
		panic(err)
	}
}
