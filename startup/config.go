package startup

import (
	"log"

	"gitee.com/kelvins-io/kelvins/config"
	"gitee.com/kelvins-io/kelvins/config/setting"
	"github.com/yongcycchen/mall-users/vars"
)

const (
	SectionEmailConfig             = "email-config"
	SectionQueueUserRegisterNotice = "queue-user-register-notice"
	SectionQueueUserStateNotice    = "queue-user-state-notice"
	SectionJwt                     = "kelvins-jwt"
)

func LoadConfig() error {
	// email setting
	log.Printf("[info] Load default config %s", SectionEmailConfig)
	vars.EmailConfigSetting = new(vars.EmailConfigSettingS)
	config.MapConfig(SectionEmailConfig, vars.EmailConfigSetting)
	// User Register Notice
	log.Printf("[info] Load default config %s", SectionQueueUserRegisterNotice)
	vars.QueueAMQPSettingUserRegisterNotice = new(setting.QueueAMQPSettingS)
	config.MapConfig(SectionQueueUserRegisterNotice, vars.QueueAMQPSettingUserRegisterNotice)
	// User State Notice
	log.Printf("[info] Load default config %s", SectionQueueUserStateNotice)
	vars.QueueAMQPSettingUserStateNotice = new(setting.QueueAMQPSettingS)
	config.MapConfig(SectionQueueUserStateNotice, vars.QueueAMQPSettingUserStateNotice)
	// User Token
	log.Printf("[info] Load default config %s", SectionJwt)
	vars.JwtSetting = new(vars.JwtSettingS)
	config.MapConfig(SectionJwt, vars.JwtSetting)
	return nil
}
