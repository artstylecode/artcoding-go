package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"strconv"
)
// FailOnError 错误时记录日志
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
// SysConfig ini配置文件类
type SysConfig struct {
	allConfig map[string]map[string]string
}
//Load 加载配置文件
func (s *SysConfig) Load(path string)  {
	config,err := ini.Load(path)
	if err != nil {
		panic("加载配置文件失败!")
	}
	configSections := config.Sections()
	fmt.Println(configSections)
	var res = make(map[string]map[string]string)
	for _,configSection := range configSections{
		sectionName := configSection.Name()
		res[sectionName] = make(map[string]string)
		for key,value := range configSection.KeysHash(){
			res[sectionName][key] = value
		}
	}
	s.allConfig = res
}
// checkIsInit 验证是否初始化
func (s *SysConfig) checkIsInit()  {
	if s.allConfig == nil{
		panic("文件未加载或文件加载失败！")
	}
}
// GetSectionConfig 获取指定段落配置
func (s SysConfig) GetSectionConfig(sectionName string)map[string]string  {
	s.checkIsInit()
	return s.allConfig[sectionName]
}
// GetAllConfig 获取所有段落配置
func (s SysConfig) GetAllConfig()map[string]map[string]string  {
	s.checkIsInit()
	return s.allConfig
}
//GetStringValue 获取string配置值
func (s SysConfig) GetStringValue(sectionName string, key string)string  {
	s.checkIsInit()
	return s.allConfig[sectionName][key]
}
// GetIntValue 获取int配置值
func (s SysConfig) GetIntValue(sectionName string, key string) int {
	s.checkIsInit()
	strval := s.allConfig[sectionName][key]
	val, err := strconv.Atoi(strval)
	if err != nil {
		panic(fmt.Sprintf("模块：%s键：%s数据转换int失败", sectionName, key))
	}
	return val
}
// GetBooleanValue 获取boolean配置值
func (s SysConfig) GetBooleanValue(sectionName string, key string) bool {
	s.checkIsInit()
	strval := s.allConfig[sectionName][key]
	val, err := strconv.ParseBool(strval)
	if err != nil {
		panic(fmt.Sprintf("模块：%s键：%s数据转换boolean失败", sectionName, key))
	}
	return val
}
// GetFloatValue 获取float配置值
func (s SysConfig) GetFloatValue(sectionName string, key string, bitSize int) float64 {
	s.checkIsInit()
	strval := s.allConfig[sectionName][key]
	val, err := strconv.ParseFloat(strval,bitSize)
	if err != nil {
		panic(fmt.Sprintf("模块：%s键：%s数据转换float失败", sectionName, key))
	}
	return val
}