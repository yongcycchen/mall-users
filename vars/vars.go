package vars

import (
	"gitee.com/kelvins-io/common/queue"
	"gitee.com/kelvins-io/kelvins/config/setting"
	"github.com/yongcycchen/mall-api/pkg/util/goroutine"
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
