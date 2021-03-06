package server

import (
	context "context"

	"gitee.com/kelvins-io/common/errcode"
	"github.com/yongcycchen/mall-users/pkg/code"
	"github.com/yongcycchen/mall-users/pkg/util"
	"github.com/yongcycchen/mall-users/proto/mall_users_proto/users"
	"github.com/yongcycchen/mall-users/service"
)

type UsersServer struct {
	users.UnimplementedUsersServiceServer
}

func NewUsersServer() users.UsersServiceServer {
	return new(UsersServer)
}

func (u *UsersServer) GetUserInfo(ctx context.Context, in *users.GetUserInfoRequest) (*users.GetUserInfoResponse, error) {
	if in.Uid <= 0 {
		return &users.GetUserInfoResponse{
			Common: &users.CommonResponse{
				Code: users.RetCode_USER_NOT_EXIST,
				Msg:  errcode.GetErrMsg(code.UserNotExist),
			},
		}, nil
	}
	userInfo, retCode := service.GetUserInfo(ctx, int(in.Uid))
	if retCode != code.Success {
		return &users.GetUserInfoResponse{
			Common: &users.CommonResponse{
				Code: users.RetCode_ERROR,
				Msg:  errcode.GetErrMsg(code.ErrorServer),
			},
		}, nil
	}
	return &users.GetUserInfoResponse{
		Common: &users.CommonResponse{
			Code: users.RetCode_SUCCESS,
			Msg:  errcode.GetErrMsg(code.Success),
		},
		Info: &users.UserInfo{
			Uid:         int64(userInfo.Id),
			AccountId:   userInfo.AccountId,
			UserName:    userInfo.UserName,
			Sex:         int32(userInfo.Sex),
			CountryCode: userInfo.CountryCode,
			Phone:       userInfo.Phone,
			Email:       userInfo.Email,
			State:       int32(userInfo.State),
			IdCardNo:    userInfo.IdCardNo.String,
			Inviter:     int64(userInfo.Inviter),
			InviterCode: userInfo.InviteCode,
			ContactAddr: userInfo.ContactAddr,
			Age:         int32(userInfo.Age),
			CreateTime:  util.ParseTimeOfStr(userInfo.CreateTime.Unix()),
			UpdateTime:  util.ParseTimeOfStr(userInfo.UpdateTime.Unix()),
		},
	}, nil
}

func (u *UsersServer) GetUserInfoByPhone(ctx context.Context, in *users.GetUserInfoByPhoneRequest) (*users.GetUserInfoByPhoneResponse, error) {
	if in.Phone == "" || in.CountryCode == "" {
		return &users.GetUserInfoByPhoneResponse{
			Common: &users.CommonResponse{
				Code: users.RetCode_USER_NOT_EXIST,
				Msg:  errcode.GetErrMsg(code.UserNotExist),
			},
		}, nil
	}
	userInfo, retCode := service.GetUserInfoByPhone(ctx, in.CountryCode, in.Phone)
	if retCode != code.Success {
		return &users.GetUserInfoByPhoneResponse{
			Common: &users.CommonResponse{
				Code: users.RetCode_ERROR,
				Msg:  errcode.GetErrMsg(code.ErrorServer),
			},
		}, nil
	}
	return &users.GetUserInfoByPhoneResponse{
		Common: &users.CommonResponse{
			Code: users.RetCode_SUCCESS,
			Msg:  errcode.GetErrMsg(code.Success),
		},
		Info: &users.UserInfo{
			Uid:         int64(userInfo.Id),
			AccountId:   userInfo.AccountId,
			UserName:    userInfo.UserName,
			Sex:         int32(userInfo.Sex),
			CountryCode: userInfo.CountryCode,
			Phone:       userInfo.Phone,
			Email:       userInfo.Email,
			State:       int32(userInfo.State),
			IdCardNo:    userInfo.IdCardNo.String,
			Inviter:     int64(userInfo.Inviter),
			InviterCode: userInfo.InviteCode,
			ContactAddr: userInfo.ContactAddr,
			Age:         int32(userInfo.Age),
			CreateTime:  util.ParseTimeOfStr(userInfo.CreateTime.Unix()),
			UpdateTime:  util.ParseTimeOfStr(userInfo.UpdateTime.Unix()),
		},
	}, nil
}

func (u *UsersServer) CheckUserByPhone(ctx context.Context, in *users.CheckUserByPhoneRequest) (*users.CheckUserByPhoneResponse, error) {
	result := &users.CheckUserByPhoneResponse{
		Common: &users.CommonResponse{
			Code: users.RetCode_SUCCESS,
			Msg:  "",
		},
		IsExist: false,
	}
	exist, retCode := service.CheckUserExist(ctx, in.CountryCode, in.Phone)
	if retCode != code.Success {
		switch retCode {
		case code.UserNotExist:
			result.Common.Code = users.RetCode_USER_NOT_EXIST
		default:
			result.Common.Code = users.RetCode_ERROR
		}
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return result, nil
	}
	result.IsExist = exist
	return result, nil
}

func (u *UsersServer) Register(ctx context.Context, in *users.RegisterRequest) (*users.RegisterResponse, error) {
	result := &users.RegisterResponse{
		Common: &users.CommonResponse{
			Code: users.RetCode_SUCCESS,
		},
	}
	reg, retCode := service.RegisterUser(ctx, in)
	if retCode != code.Success {
		switch retCode {
		case code.UserExist:
			result.Common.Code = users.RetCode_USER_EXIST
		default:
			result.Common.Code = users.RetCode_ERROR
		}
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return result, nil
	}
	result.Result = &users.RegisterResult{InviteCode: reg.InviteCode}
	return result, nil
}

func (u *UsersServer) LoginUser(ctx context.Context, in *users.LoginUserRequest) (*users.LoginUserResponse, error) {
	result := &users.LoginUserResponse{
		Common: &users.CommonResponse{
			Code: users.RetCode_SUCCESS,
		},
	}
	token, retCode := service.LoginUser(ctx, in)
	if retCode != code.Success {
		switch retCode {
		case code.UserNotExist:
			result.Common.Code = users.RetCode_USER_NOT_EXIST
		case code.UserPwdNotMatch:
			result.Common.Code = users.RetCode_USER_PWD_NOT_MATCH
		default:
			result.Common.Code = users.RetCode_ERROR
		}
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return result, nil
	}
	result.IdentityToken = token
	return result, nil
}

func (u *UsersServer) PasswordReset(ctx context.Context, in *users.PasswordResetRequest) (*users.PasswordResetResponse, error) {
	result := &users.PasswordResetResponse{Common: &users.CommonResponse{
		Code: users.RetCode_SUCCESS,
		Msg:  "",
	}}
	retCode := service.PasswordReset(ctx, in)
	if retCode != code.Success {
		switch retCode {
		case code.UserNotExist:
			result.Common.Code = users.RetCode_USER_NOT_EXIST
		default:
			result.Common.Code = users.RetCode_ERROR
		}
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return result, nil
	}
	return result, nil
}

func (u *UsersServer) CheckUserIdentity(ctx context.Context, in *users.CheckUserIdentityRequest) (*users.CheckUserIdentityResponse, error) {
	result := &users.CheckUserIdentityResponse{Common: &users.CommonResponse{
		Code: users.RetCode_SUCCESS,
		Msg:  "",
	}}
	retCode := service.CheckUserIdentity(ctx, in)
	if retCode != code.Success {
		switch retCode {
		case code.UserNotExist:
			result.Common.Code = users.RetCode_USER_NOT_EXIST
		case code.UserPwdNotMatch:
			result.Common.Code = users.RetCode_USER_PWD_NOT_MATCH
		default:
			result.Common.Code = users.RetCode_ERROR
		}
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return result, nil
	}
	return result, nil
}

func (u *UsersServer) UpdateUserLoginState(ctx context.Context, in *users.UpdateUserLoginStateRequest) (*users.UpdateUserLoginStateResponse, error) {
	result := &users.UpdateUserLoginStateResponse{Common: &users.CommonResponse{
		Code: users.RetCode_SUCCESS,
		Msg:  "",
	}}
	retCode := service.UpdateUserLoginState(ctx, in)
	if retCode != code.Success {
		switch retCode {
		case code.UserNotExist:
			result.Common.Code = users.RetCode_USER_NOT_EXIST
		default:
			result.Common.Code = users.RetCode_ERROR
		}
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return result, nil
	}
	return result, nil
}

func (u *UsersServer) GetUserInfoByInviteCode(ctx context.Context, in *users.GetUserByInviteCodeRequest) (*users.GetUserByInviteCodeResponse, error) {
	userInfo, retCode := service.GetUserInfoByInviteCode(ctx, in.InviteCode)
	if retCode != code.Success {
		return &users.GetUserByInviteCodeResponse{
			Common: &users.CommonResponse{
				Code: users.RetCode_ERROR,
				Msg:  errcode.GetErrMsg(code.ErrorServer),
			},
		}, nil
	}
	return &users.GetUserByInviteCodeResponse{
		Common: &users.CommonResponse{
			Code: users.RetCode_SUCCESS,
			Msg:  errcode.GetErrMsg(code.Success),
		},
		Info: &users.UserInfo{
			Uid:         int64(userInfo.Id),
			AccountId:   userInfo.AccountId,
			UserName:    userInfo.UserName,
			Sex:         int32(userInfo.Sex),
			CountryCode: userInfo.CountryCode,
			Phone:       userInfo.Phone,
			Email:       userInfo.Email,
			State:       int32(userInfo.State),
			IdCardNo:    userInfo.IdCardNo.String,
			Inviter:     int64(userInfo.Inviter),
			InviterCode: userInfo.InviteCode,
			ContactAddr: userInfo.ContactAddr,
			Age:         int32(userInfo.Age),
			CreateTime:  util.ParseTimeOfStr(userInfo.CreateTime.Unix()),
			UpdateTime:  util.ParseTimeOfStr(userInfo.UpdateTime.Unix()),
		},
	}, nil
}

func (u *UsersServer) ModifyUserDeliveryInfo(ctx context.Context, in *users.ModifyUserDeliveryInfoRequest) (*users.ModifyUserDeliveryInfoResponse, error) {
	result := users.ModifyUserDeliveryInfoResponse{Common: &users.CommonResponse{
		Code: users.RetCode_SUCCESS,
		Msg:  "",
	}}
	retCode := service.ModifyUserDeliveryInfo(ctx, in)
	if retCode != code.Success {
		switch retCode {
		case code.TransactionFailed:
			result.Common.Code = users.RetCode_TRANSACTION_FAILED
		case code.UserDeliveryInfoExist:
			result.Common.Code = users.RetCode_USER_DELIVERY_INFO_EXIST
		case code.UserDeliveryInfoNotExist:
			result.Common.Code = users.RetCode_USER_DELIVERY_INFO_NOT_EXIST
		case code.ErrorServer:
			result.Common.Code = users.RetCode_ERROR
		}
		return &result, nil
	}
	return &result, nil
}

func (u *UsersServer) GetUserDeliveryInfo(ctx context.Context, in *users.GetUserDeliveryInfoRequest) (*users.GetUserDeliveryInfoResponse, error) {
	result := &users.GetUserDeliveryInfoResponse{Common: &users.CommonResponse{
		Code: users.RetCode_SUCCESS,
	}, Info: make([]*users.UserDeliveryInfo, 0)}
	list, retCode := service.GetUserDeliveryInfo(ctx, in)
	if retCode != code.Success {
		result.Common.Code = users.RetCode_ERROR
		return result, nil
	}
	result.Info = list
	return result, nil
}

func (u *UsersServer) FindUserInfo(ctx context.Context, in *users.FindUserInfoRequest) (*users.FindUserInfoResponse, error) {
	result := &users.FindUserInfoResponse{
		Common: &users.CommonResponse{
			Code: users.RetCode_SUCCESS,
		},
		InfoList: nil,
	}
	userInfoList, retCode := service.FindUserInfo(ctx, in)
	if retCode != code.Success {
		result.Common.Code = users.RetCode_ERROR
		return result, nil
	}
	result.InfoList = userInfoList
	return result, nil
}

func (u *UsersServer) UserAccountCharge(ctx context.Context, in *users.UserAccountChargeRequest) (*users.UserAccountChargeResponse, error) {
	result := &users.UserAccountChargeResponse{Common: &users.CommonResponse{
		Code: users.RetCode_SUCCESS,
	}}
	retCode := service.UserAccountCharge(ctx, in)
	if retCode != code.Success {
		switch retCode {
		case code.UserNotExist:
			result.Common.Code = users.RetCode_USER_NOT_EXIST
		case code.AccountNotExist:
			result.Common.Code = users.RetCode_ACCOUNT_NOT_EXIST
		case code.TransactionFailed:
			result.Common.Code = users.RetCode_TRANSACTION_FAILED
		case code.AccountStateLock:
			result.Common.Code = users.RetCode_ACCOUNT_LOCK
		case code.AccountStateInvalid:
			result.Common.Code = users.RetCode_ACCOUNT_INVALID
		case code.UserChargeRun:
			result.Common.Code = users.RetCode_USER_CHARGE_RUN
		case code.UserChargeSuccess:
			result.Common.Code = users.RetCode_USER_CHARGE_SUCCESS
		case code.UserChargeTradeNoEmpty:
			result.Common.Code = users.RetCode_USER_CHARGE_TRADE_NO_EMPTY
		default:
			result.Common.Code = users.RetCode_ERROR
		}
		return result, nil
	}
	return result, nil
}

func (u *UsersServer) CheckUserDeliveryInfo(ctx context.Context, in *users.CheckUserDeliveryInfoRequest) (*users.CheckUserDeliveryInfoResponse, error) {
	result := &users.CheckUserDeliveryInfoResponse{Common: &users.CommonResponse{
		Code: users.RetCode_SUCCESS,
	}}
	retCode := service.CheckUserDeliveryInfo(ctx, in)
	if retCode != code.Success {
		switch retCode {
		case code.UserDeliveryInfoNotExist:
			result.Common.Code = users.RetCode_USER_DELIVERY_INFO_NOT_EXIST
		default:
			result.Common.Code = users.RetCode_ERROR
		}
		return result, nil
	}
	return result, nil
}

func (u *UsersServer) CheckUserState(ctx context.Context, in *users.CheckUserStateRequest) (*users.CheckUserStateResponse, error) {
	result := &users.CheckUserStateResponse{Common: &users.CommonResponse{
		Code: users.RetCode_SUCCESS,
	}}
	retCode := service.CheckUserState(ctx, in)
	if retCode != code.Success {
		switch retCode {
		case code.UserNotExist:
			result.Common.Code = users.RetCode_USER_NOT_EXIST
		default:
			result.Common.Code = users.RetCode_ERROR
		}
		return result, nil
	}
	return result, nil
}

func (u *UsersServer) GetUserAccountId(ctx context.Context, in *users.GetUserAccountIdRequest) (*users.GetUserAccountIdResponse, error) {
	result := &users.GetUserAccountIdResponse{
		Common: &users.CommonResponse{
			Code: users.RetCode_SUCCESS,
			Msg:  "",
		},
		InfoList: nil,
	}
	accountInfoList, retCode := service.GetUserAccountId(ctx, in)
	if retCode != code.Success {
		switch retCode {
		case code.UserNotExist:
			result.Common.Code = users.RetCode_USER_NOT_EXIST
		default:
			result.Common.Code = users.RetCode_ERROR
		}
		return result, nil
	}
	result.InfoList = accountInfoList
	return result, nil
}

func (u *UsersServer) ListUserInfo(ctx context.Context, in *users.ListUserInfoRequest) (*users.ListUserInfoResponse, error) {
	result := &users.ListUserInfoResponse{
		Common: &users.CommonResponse{
			Code: users.RetCode_SUCCESS,
		},
		UserInfoList: nil,
	}
	userInfoList, retCode := service.ListUserInfo(ctx, in)
	result.UserInfoList = userInfoList
	if retCode != code.Success {
		result.Common.Code = users.RetCode_ERROR
		return result, nil
	}
	return result, nil
}
