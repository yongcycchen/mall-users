package vars

import (
	"github.com/yongcycchen/mall-api/pkg/util/goroutine"
	"gitee.com/kelvins-io/common/queue"
	"gitee.com/kelvins-io/kelvins/config/setting"
)

var (
	App                                *GRPCApplication
	EmailConfigSetting                 *EmailConfigSettingS
	JwtSetting                         *JwtSettingS
	QueueAMQPSettingUserRegisterNotice *setting.QueueAMQPSettingS
	QueueServerUserRegisterNotice      *queue.MachineryQueue
	QueueAMQPSettingUserStateNotice    *setting.QueueAMQPSettingS
	QueueServerUserStateNotice         *queue.MachineryQueue
	GPool                              *goroutine.Pool
)
