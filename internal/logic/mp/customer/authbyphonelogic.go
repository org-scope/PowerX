package customer

import (
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"PowerX/internal/types/models"
	"PowerX/internal/uc/powerx"
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthByPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthByPhoneLogic {
	return &AuthByPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthByPhoneLogic) AuthByPhone(req *types.MPCustomerAuthRequest) (resp *types.MPCustomerLoginAuthReply, err error) {
	//fmt.DD(req)

	// 获取session数据
	rs, err := l.svcCtx.PowerX.WechatMP.App.Auth.Session(l.ctx, req.Code)
	if err != nil {
		panic(err)
		return
	}
	//rs := &response.ResponseCode2Session{
	//	OpenID:     "o1IFX5A8sfi5nbkXwOzNLLLiL0OA",
	//	SessionKey: "rUoiNCDNWekX68d7TmnNGw==",
	//}

	//fmt.Dump(rs, req)
	// 解码手机授权信息
	msgData, errEncrypt := l.svcCtx.PowerX.WechatMP.App.Encryptor.DecryptData(req.EncryptedData, rs.SessionKey, req.IV)

	if errEncrypt != nil {
		panic(errEncrypt.ErrMsg)
		return
	}

	println(string(msgData))
	// 解析手机信息
	mpPhoneInfo := &models.MPPhoneInfo{}
	err = object.JsonDecode(msgData, mpPhoneInfo)
	if err != nil {
		panic(err.Error())
		return
	}

	mpCustomer := &models.WechatMPCustomer{
		OpenID:     rs.OpenID,
		SessionKey: rs.SessionKey,
		UnionID:    rs.UnionID,
	}

	// 转化手机信息给小程序客户记录
	mpCustomer = l.svcCtx.PowerX.WechatMP.ConvertMPPhoneInfoToMPCustomer(l.ctx, mpPhoneInfo, mpCustomer)

	// upsert 小程序客户记录
	mpCustomer, err = l.svcCtx.PowerX.WechatMP.UpsertMPCustomer(l.ctx, mpCustomer)
	if err != nil {
		panic(err)
		return
	}

	token := l.svcCtx.PowerX.CustomerAuthorization.SignToken(mpCustomer, l.svcCtx.Config.JWTSecret)

	return &types.MPCustomerLoginAuthReply{
		OpenID:      mpCustomer.OpenID,
		UnionID:     mpCustomer.UnionID,
		PhoneNumber: mpCustomer.PhoneNumber,
		NickName:    mpCustomer.NickName,
		AvatarURL:   mpCustomer.AvatarURL,
		Gender:      mpCustomer.Gender,
		Token: types.Token{
			TokenType:    token.TokenType,
			ExpiresIn:    fmt.Sprintf("%d", powerx.CustomerTokenExpiredDuration),
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		},
	}, nil
}