# Access token 获取

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execGetAccessToken`|`reqAccessToken`|`respAccessToken`|-|`GET /cgi-bin/gettoken`|[获取access_token](https://work.weixin.qq.com/api/doc#90000/90135/91039)
`execGetJSAPITicket`|`reqJSAPITicket`|`respJSAPITicket`|+|`GET /cgi-bin/get_jsapi_ticket`|[获取企业的jsapi_ticket](https://open.work.weixin.qq.com/api/doc/90000/90136/90506)
`execGetJSAPITicketAgentConfig`|`reqJSAPITicketAgentConfig`|`respJSAPITicket`|+|`GET /cgi-bin/ticket/get`|[获取应用的jsapi_ticket](https://open.work.weixin.qq.com/api/doc/90000/90136/90506)
`execJSCode2Session`|`reqJSCode2Session`|`respJSCode2Session`|+|`GET /cgi-bin/miniprogram/jscode2session`|[临时登录凭证校验code2Session](https://open.work.weixin.qq.com/api/doc/90000/90136/91507)
`execAuthCode2UserInfo`|`reqAuthCode2UserInfo`|`respAuthCode2UserInfo`|+|`GET /cgi-bin/auth/getuserinfo`|[获取访问用户身份](https://developer.work.weixin.qq.com/document/path/91023)

# 成员管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execUserCreate`|TODO|TODO|+|`POST /cgi-bin/user/create`|[创建成员](https://work.weixin.qq.com/api/doc#90000/90135/90195)
`execUserGet`|`reqUserGet`|`respUserGet`|+|`GET /cgi-bin/user/get`|[读取成员](https://work.weixin.qq.com/api/doc#90000/90135/90196)
`execUserUpdate`|`reqUserUpdate`|`respUserUpdate`|+|`POST /cgi-bin/user/update`|[更新成员](https://work.weixin.qq.com/api/doc#90000/90135/90197)
`execUserDelete`|TODO|TODO|+|`GET /cgi-bin/user/delete`|[删除成员](https://work.weixin.qq.com/api/doc#90000/90135/90198)
`execUserBatchDelete`|TODO|TODO|+|`POST /cgi-bin/user/batchdelete`|[批量删除成员](https://work.weixin.qq.com/api/doc#90000/90135/90199)
`execUserSimpleList`|TODO|TODO|+|`GET /cgi-bin/user/simplelist`|[获取部门成员](https://work.weixin.qq.com/api/doc#90000/90135/90200)
`execUserList`|`reqUserList`|`respUserList`|+|`GET /cgi-bin/user/list`|[获取部门成员详情](https://work.weixin.qq.com/api/doc#90000/90135/90201)
`execConvertUserIDToOpenID`|`reqConvertUserIDToOpenID`|`respConvertUserIDToOpenID`|+|`POST /cgi-bin/user/convert_to_openid`|[userid转openid](https://work.weixin.qq.com/api/doc#90000/90135/90202)
`execConvertOpenIDToUserID`|`reqConvertOpenIDToUserID`|`respConvertOpenIDToUserID`|+|`POST /cgi-bin/user/convert_to_userid`|[openid转userid](https://work.weixin.qq.com/api/doc#90000/90135/90202)
`execUserAuthSucc`|TODO|TODO|+|`GET /cgi-bin/user/authsucc`|[二次验证](https://work.weixin.qq.com/api/doc#90000/90135/90203)
`execUserBatchInvite`|TODO|TODO|+|`POST /cgi-bin/batch/invite`|[邀请成员](https://work.weixin.qq.com/api/doc#90000/90135/90975)
`execUserJoinQrcode`|`reqUserJoinQrcode`|`respUserJoinQrcode`|+|`GET /cgi-bin/corp/get_join_qrcode`|[获取加入企业二维码](https://developer.work.weixin.qq.com/document/path/91714)
`execUserIDByMobile`|`reqUserIDByMobile`|`respUserIDByMobile`|+|`POST /cgi-bin/user/getuserid`|[手机号获取userid](https://work.weixin.qq.com/api/doc/90001/90143/91693)
`execUserIDByEmail`|`reqUserIDByEmail`|`respUserIDByEmail`|+|`POST /cgi-bin/user/get_userid_by_email`|[邮箱获取userid](https://developer.work.weixin.qq.com/document/path/95895)

# 部门管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execDeptCreate`|`reqDeptCreate`|`respDeptCreate`|+|`POST /cgi-bin/department/create`|[创建部门](https://work.weixin.qq.com/api/doc#90000/90135/90205)
`execDeptUpdate`|TODO|TODO|+|`POST /cgi-bin/department/update`|[更新部门](https://work.weixin.qq.com/api/doc#90000/90135/90206)
`execDeptDelete`|TODO|TODO|+|`GET /cgi-bin/department/delete`|[删除部门](https://work.weixin.qq.com/api/doc#90000/90135/90207)
`execDeptList`|`reqDeptList`|`respDeptList`|+|`GET /cgi-bin/department/list`|[获取部门列表](https://work.weixin.qq.com/api/doc#90000/90135/90208)
`execDeptSimpleList`|`reqDeptSimpleList`| `respDeptSimpleList` |+|`GET /cgi-bin/department/simplelist`|[获取子部门ID列表](https://developer.work.weixin.qq.com/document/path/95350)

# 标签管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execTagCreate`|TODO|TODO|+|`POST /cgi-bin/tag/create`|[创建标签](https://work.weixin.qq.com/api/doc#90000/90135/90210)
`execTagUpdate`|TODO|TODO|+|`POST /cgi-bin/tag/update`|[更新标签名字](https://work.weixin.qq.com/api/doc#90000/90135/90211)
`execTagDelete`|TODO|TODO|+|`GET /cgi-bin/tag/delete`|[删除标签](https://work.weixin.qq.com/api/doc#90000/90135/90212)
`execTagListUsers`|TODO|TODO|+|`GET /cgi-bin/tag/get`|[获取标签成员](https://work.weixin.qq.com/api/doc#90000/90135/90213)
`execTagAddUsers`|TODO|TODO|+|`POST /cgi-bin/tag/addtagusers`|[增加标签成员](https://work.weixin.qq.com/api/doc#90000/90135/90214)
`execTagDeleteUsers`|TODO|TODO|+|`POST /cgi-bin/tag/deltagusers`|[删除标签成员](https://work.weixin.qq.com/api/doc#90000/90135/90215)
`execTagList`|TODO|TODO|+|`GET /cgi-bin/tag/list`|[获取标签列表](https://work.weixin.qq.com/api/doc#90000/90135/90216)

# 异步批量接口

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--

# 身份验证

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execUserInfoGet`|`reqUserInfoGet`|`respUserInfoGet`|+|`GET /cgi-bin/user/getuserinfo`|[获取访问用户身份](https://work.weixin.qq.com/api/doc/90000/90135/91023)

# 外部联系人管理 - 客户管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execExternalContactList`|`reqExternalContactList`|`respExternalContactList`|+|`GET /cgi-bin/externalcontact/list`|[获取客户列表](https://work.weixin.qq.com/api/doc/90000/90135/92113)
`execExternalContactGet`|`reqExternalContactGet`|`respExternalContactGet`|+|`GET /cgi-bin/externalcontact/get`|[获取客户详情](https://work.weixin.qq.com/api/doc/90000/90135/92114)
`execExternalContactBatchList`|`reqExternalContactBatchList`|`respExternalContactBatchList`|+|`POST /cgi-bin/externalcontact/batch/get_by_user`|[批量获取客户详情](https://work.weixin.qq.com/api/doc/90000/90135/92994)
`execExternalContactRemark`|`reqExternalContactRemark`|`respExternalContactRemark`|+|`POST /cgi-bin/externalcontact/remark`|[修改客户备注信息](https://work.weixin.qq.com/api/doc/90000/90135/92115)

# 外部联系人管理 - 客户标签管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execExternalContactListCorpTags`|`reqExternalContactListCorpTags`|`respExternalContactListCorpTags`|+|`POST /cgi-bin/externalcontact/get_corp_tag_list`|[获取企业标签库](https://work.weixin.qq.com/api/doc/90000/90135/92117)
`execExternalContactAddCorpTag`|`reqExternalContactAddCorpTagGroup`|`respExternalContactAddCorpTag`|+|`POST /cgi-bin/externalcontact/add_corp_tag`|[添加企业客户标签](https://work.weixin.qq.com/api/doc/90000/90135/92117)
`execExternalContactEditCorpTag`|`reqExternalContactEditCorpTag`|`respExternalContactEditCorpTag`|+|`POST /cgi-bin/externalcontact/edit_corp_tag`|[编辑企业客户标签](https://work.weixin.qq.com/api/doc/90000/90135/92117)
`execExternalContactDelCorpTag`|`reqExternalContactDelCorpTag`|`respExternalContactDelCorpTag`|+|`POST /cgi-bin/externalcontact/del_corp_tag`|[删除企业客户标签](https://work.weixin.qq.com/api/doc/90000/90135/92117)
`execExternalContactMarkTag`|`reqExternalContactMarkTag`|`respExternalContactMarkTag`|+|`POST /cgi-bin/externalcontact/mark_tag`|[标记客户企业标签](https://work.weixin.qq.com/api/doc/90000/90135/92118)

# 外部联系人管理 - 客户分配

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execListUnassignedExternalContact`|`reqListUnassignedExternalContact`|`respListUnassignedExternalContact`|+|`POST /cgi-bin/externalcontact/get_unassigned_list`|[获取离职成员的客户列表](https://work.weixin.qq.com/api/doc/90000/90135/92124)
`execTransferExternalContact`|`reqTransferExternalContact`|`respTransferExternalContact`|+|`POST /cgi-bin/externalcontact/transfer`|[分配成员的客户](https://work.weixin.qq.com/api/doc/90000/90135/92125)
`execGetTransferExternalContactResult`|`reqGetTransferExternalContactResult`|`respGetTransferExternalContactResult`|+|`POST /cgi-bin/externalcontact/get_transfer_result`|[查询客户接替结果](https://work.weixin.qq.com/api/doc/90000/90135/92973)
`execTransferGroupChatExternalContact`|`reqTransferGroupChatExternalContact`|`respTransferGroupChatExternalContact`|+|`POST /cgi-bin/externalcontact/groupchat/transfer`|[离职成员的群再分配](https://work.weixin.qq.com/api/doc/90000/90135/92127)

# 应用管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execAgentGet`|TODO|TODO|+|`GET /cgi-bin/agent/get`|[获取指定的应用详情](https://work.weixin.qq.com/api/doc#90000/90135/90227)
`execAgentList`|TODO|TODO|+|`GET /cgi-bin/agent/list`|[获取access_token对应的应用列表](https://work.weixin.qq.com/api/doc#90000/90135/90227)
`execAgentSet`|TODO|TODO|+|`POST /cgi-bin/agent/set`|[设置应用](https://work.weixin.qq.com/api/doc#90000/90135/90228)

# 应用管理 - 自定义菜单

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execMenuCreate`|TODO|TODO|+|`POST /cgi-bin/menu/create`|[创建菜单](https://work.weixin.qq.com/api/doc#90000/90135/90231)
`execMenuGet`|TODO|TODO|+|`GET /cgi-bin/menu/get`|[获取菜单](https://work.weixin.qq.com/api/doc#90000/90135/90232)
`execMenuDelete`|TODO|TODO|+|`GET /cgi-bin/menu/delete`|[删除菜单](https://work.weixin.qq.com/api/doc#90000/90135/90233)

# 消息推送

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execAppchatCreate`|`reqAppchatCreate`|`respAppchatCreate`|+|`POST /cgi-bin/appchat/create`|[创建群聊会话](https://work.weixin.qq.com/api/doc#90000/90135/90245)
`execAppchatUpdate`|`reqAppchatUpdate`|`respAppchatUpdate`|+|`POST /cgi-bin/appchat/update`|[修改群聊会话](https://work.weixin.qq.com/api/doc#90000/90135/90246)
`execAppchatGet`|`reqAppchatGet`|`respAppchatGet`|+|`GET /cgi-bin/appchat/get`|[获取群聊会话](https://work.weixin.qq.com/api/doc#90000/90135/90247)
`execMessageSend`|`reqMessage`|`respMessageSend`|+|`POST /cgi-bin/message/send`|[发送应用消息](https://work.weixin.qq.com/api/doc#90000/90135/90236)
`execAppchatSend`|`reqMessage`|`respMessageSend`|+|`POST /cgi-bin/appchat/send`|[应用推送消息](https://work.weixin.qq.com/api/doc#90000/90135/90248)

# 素材管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execMediaUpload`|`reqMediaUpload`|`respMediaUpload`|+|`POST(media) /cgi-bin/media/upload`|[上传临时素材](https://work.weixin.qq.com/api/doc#90000/90135/90253)
`execMediaUploadImg`|`reqMediaUploadImg`|`respMediaUploadImg`|+|`POST(media) /cgi-bin/media/uploadimg`|[上传永久图片](https://work.weixin.qq.com/api/doc#90000/90135/90256)
`execMediaGet`|TODO|TODO|+|`GET /cgi-bin/media/get`|[获取临时素材](https://work.weixin.qq.com/api/doc#90000/90135/90254)
`execMediaGetJSSDK`|TODO|TODO|+|`GET /cgi-bin/media/get/jssdk`|[获取高清语音素材](https://work.weixin.qq.com/api/doc#90000/90135/90255)

# OA 数据接口

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execOAGetTemplateDetail`|`reqOAGetTemplateDetail`|`respOAGetTemplateDetail`|+|`POST /cgi-bin/oa/gettemplatedetail`|[获取审批模板详情](https://work.weixin.qq.com/api/doc/90000/90135/91982)
`execOAApplyEvent`|`reqOAApplyEvent`|`respOAApplyEvent`|+|`POST /cgi-bin/oa/applyevent`|[提交审批申请](https://work.weixin.qq.com/api/doc/90000/90135/91853)
`execOAGetApprovalInfo`|`reqOAGetApprovalInfo`|`respOAGetApprovalInfo`|+|`POST /cgi-bin/oa/getapprovalinfo`|[批量获取审批单号](https://work.weixin.qq.com/api/doc/90000/90135/91816)
`execOAGetApprovalDetail`|`reqOAGetApprovalDetail`|`respOAGetApprovalDetail`|+|`POST /cgi-bin/oa/getapprovaldetail`|[获取审批申请详情](https://work.weixin.qq.com/api/doc/90000/90135/91983)
`execOAGetCorpVacationConf`| `reqOAGetCorpVacationConf`     | `respOAGetCorpVacationConf`     |+| `GET /cgi-bin/oa/vacation/getcorpconf`           |[获取企业假期管理配置](https://developer.work.weixin.qq.com/document/path/93375)
`execOAGetUserVacationQuota`| `reqOAGetUserVacationQuota`    | `respOAGetUserVacationQuota`    |+| `POST /cgi-bin/oa/vacation/getuservacationquota` |[获取成员假期余额](https://developer.work.weixin.qq.com/document/path/93376)
`execOASetOneUserVacationQuota`| `reqOASetOneUserVacationQuota` | `respOASetOneUserVacationQuota` |+| `POST /cgi-bin/oa/vacation/setoneuserquota`      |[修改成员假期余额](https://developer.work.weixin.qq.com/document/path/93377)

# 企业支付

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--

# 电子发票

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--

# 会话内容存档

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execMsgAuditListPermitUser`|`reqMsgAuditListPermitUser`|`respMsgAuditListPermitUser`|+|`POST /cgi-bin/msgaudit/get_permit_user_list`|[获取会话内容存档开启成员列表](https://work.weixin.qq.com/api/doc/90000/90135/91614)
`execMsgAuditCheckSingleAgree`|`reqMsgAuditCheckSingleAgree`|`respMsgAuditCheckSingleAgree`|+|`POST /cgi-bin/msgaudit/check_single_agree`|[获取会话同意情况（单聊）](https://work.weixin.qq.com/api/doc/90000/90135/91782)
`execMsgAuditCheckRoomAgree`|`reqMsgAuditCheckRoomAgree`|`respMsgAuditCheckRoomAgree`|+|`POST /cgi-bin/msgaudit/check_room_agree`|[获取会话同意情况（群聊）](https://work.weixin.qq.com/api/doc/90000/90135/91782)
`execMsgAuditGetGroupChat`|`reqMsgAuditGetGroupChat`|`respMsgAuditGetGroupChat`|+|`POST /cgi-bin/msgaudit/groupchat/get`|[获取会话内容存档内部群信息](https://work.weixin.qq.com/api/doc/90000/90135/92951)

# 企业服务人员管理 - 联系我与客户入群方式

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execListFollowUserExternalContact`|`reqListFollowUserExternalContact`|`respListFollowUserExternalContact`|+|`GET /cgi-bin/externalcontact/get_follow_user_list`|[获取配置了客户联系功能的成员列表](https://developer.work.weixin.qq.com/document/path/92571)
`execAddContactExternalContact`|`reqAddContactExternalContact`|`respAddContactExternalContact`|+|`POST /cgi-bin/externalcontact/add_contact_way`|[配置客户联系「联系我」方式](https://developer.work.weixin.qq.com/document/path/92572#%E9%85%8D%E7%BD%AE%E5%AE%A2%E6%88%B7%E8%81%94%E7%B3%BB%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F)
`execGetContactWayExternalContact`|`reqGetContactWayExternalContact`|`respGetContactWayExternalContact`|+|`POST /cgi-bin/externalcontact/get_contact_way`|[获取企业已配置的「联系我」方式](https://developer.work.weixin.qq.com/document/path/92572#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F)
`execListContactWayChatExternalContact`|`reqListContactWayExternalContact`|`respListContactWayChatExternalContact`|+|`POST /cgi-bin/externalcontact/list_contact_way`|[获取企业已配置的「联系我」列表](https://developer.work.weixin.qq.com/document/path/92572#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E5%88%97%E8%A1%A8)
`execUpdateContactWayExternalContact`|`reqUpdateContactWayExternalContact`|`respUpdateContactWayExternalContact`|+|`POST /cgi-bin/externalcontact/update_contact_way`|[更新企业已配置的「联系我」成员配置](https://developer.work.weixin.qq.com/document/path/92572#%E6%9B%B4%E6%96%B0%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F)
`execDelContactWayExternalContact`|`reqDelContactWayExternalContact`|`respDelContactWayExternalContact`|+|`POST /cgi-bin/externalcontact/del_contact_way`|[删除企业已配置的「联系我」方式](https://developer.work.weixin.qq.com/document/path/92572#%E5%88%A0%E9%99%A4%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F)
`execCloseTempChatExternalContact`|`reqCloseTempChatExternalContact`|`respCloseTempChatExternalContact`|+|`POST /cgi-bin/externalcontact/close_temp_chat`|[结束临时会话](https://developer.work.weixin.qq.com/document/path/92572#%E7%BB%93%E6%9D%9F%E4%B8%B4%E6%97%B6%E4%BC%9A%E8%AF%9D)
`execAddGroupChatJoinWayExternalContact`|`reqAddGroupChatJoinWayExternalContact`|`respAddGroupChatJoinWayExternalContact`|+|`POST /cgi-bin/externalcontact/groupchat/add_join_way`|[配置客户群「加入群聊」方式](https://developer.work.weixin.qq.com/document/path/92229#%E9%85%8D%E7%BD%AE%E5%AE%A2%E6%88%B7%E7%BE%A4%E8%BF%9B%E7%BE%A4%E6%96%B9%E5%BC%8F)
`execGetGroupChatJoinWayExternalContact`|`reqGetGroupChatJoinWayExternalContact`|`respGetGroupChatJoinWayExternalContact`|+|`POST /cgi-bin/externalcontact/groupchat/get_join_way`|[获取企业已配置的客户群「加入群聊」方式](https://developer.work.weixin.qq.com/document/path/92229#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E7%BE%A4%E8%BF%9B%E7%BE%A4%E6%96%B9%E5%BC%8F%E9%85%8D%E7%BD%AE)
`execUpdateGroupChatJoinWayExternalContact`|`reqUpdateGroupChatJoinWayExternalContact`|`respUpdateGroupChatJoinWayExternalContact`|+|`POST /cgi-bin/externalcontact/groupchat/update_join_way`|[更新企业已配置的客户群「加入群聊」方式](https://developer.work.weixin.qq.com/document/path/92229#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E7%BE%A4%E8%BF%9B%E7%BE%A4%E6%96%B9%E5%BC%8F%E9%85%8D%E7%BD%AE)
`execDelGroupChatJoinWayExternalContact`|`reqDelGroupChatJoinWayExternalContact`|`respDelGroupChatJoinWayExternalContact`|+|`POST /cgi-bin/externalcontact/groupchat/del_join_way`|[删除企业已配置的客户群「加入群聊」方式](https://developer.work.weixin.qq.com/document/path/92229#%E5%88%A0%E9%99%A4%E5%AE%A2%E6%88%B7%E7%BE%A4%E8%BF%9B%E7%BE%A4%E6%96%B9%E5%BC%8F%E9%85%8D%E7%BD%AE)

# 客户联系 - 客户群管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execGroupChatListGet`|`reqGroupChatList`|`respGroupChatList`|+|`POST /cgi-bin/externalcontact/groupchat/list`|[获取客户群列表](https://developer.work.weixin.qq.com/document/path/92120)
`execGroupChatInfoGet`|`reqGroupChatInfo`|`respGroupChatInfo`|+|`POST /cgi-bin/externalcontact/groupchat/get`|[获取客户群详细](https://developer.work.weixin.qq.com/document/path/92122)
`execConvertOpenGIDToChatID`|`reqConvertOpenGIDToChatID`|`respConvertOpenGIDToChatID`|+|`POST /cgi-bin/externalcontact/opengid_to_chatid`|[客户群opengid转换](https://developer.work.weixin.qq.com/document/path/94822)

# 在职继承

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execTransferCustomer`|`reqTransferCustomer`|`respTransferCustomer`|+|`POST /cgi-bin/externalcontact/transfer_customer`|[在职继承 分配在职成员的客户](https://developer.work.weixin.qq.com/document/path/92125)
`execGetTransferCustomerResult`|`reqGetTransferCustomerResult`|`respGetTransferCustomerResult`|+|`POST /cgi-bin/externalcontact/transfer_result`|[在职继承 查询客户接替状态](https://developer.work.weixin.qq.com/document/path/94088)

# 离职继承

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execTransferResignedCustomer`|`reqTransferCustomer`|`respTransferCustomer`|+|`POST /cgi-bin/externalcontact/resigned/transfer_customer`|[离职继承 分配离职成员的客户](https://developer.work.weixin.qq.com/document/path/94081)
`execGetTransferResignedCustomerResult`|`reqGetTransferCustomerResult`|`respGetTransferCustomerResult`|+|`POST /cgi-bin/externalcontact/resigned/transfer_result`|[离职继承 查询客户接替状态](https://developer.work.weixin.qq.com/document/path/94082)

# 客户联系 - 消息推送

## API calls
Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execAddMsgTemplate`|`reqAddMsgTemplateExternalContact`|`respAddMsgTemplateExternalContact`|+|`POST /cgi-bin/externalcontact/add_msg_template`|[创建企业群发](https://developer.work.weixin.qq.com/document/path/92135)
`execSendWelcomeMsg`|`reqSendWelcomeMsgExternalContact`|`respSendWelcomeMsgExternalContact`|+|`POST /cgi-bin/externalcontact/send_welcome_msg`|[发送新客户欢迎语](https://developer.work.weixin.qq.com/document/path/92137)


# 微信客服 - 客服账号管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execKfAccountCreate`|`reqKfAccountCreate`|`respKfAccountCreate`|+|`POST /cgi-bin/kf/account/add`|[添加客服账号](https://developer.work.weixin.qq.com/document/path/94662)
`execKfAccountUpdate`|`reqKfAccountUpdate`|`respKfAccountUpdate`|+|`POST /cgi-bin/kf/account/update`|[修改客服账号](https://developer.work.weixin.qq.com/document/path/94664)
`execKfAccountDelete`|`reqKfAccountDelete`|`respKfAccountDelete`|+|`POST /cgi-bin/kf/account/del`|[删除客服账号](https://developer.work.weixin.qq.com/document/path/94663)
`execKfAccountList`|`reqKfAccountList`|`respKfAccountList`|+|`GET /cgi-bin/kf/account/list`|[获取客服账号列表](https://developer.work.weixin.qq.com/document/path/94661)
`execAddKfContact`|`reqAddKfContact`|`respAddKfContact`|+|`POST /cgi-bin/kf/add_contact_way`|[获取客服账号链接](https://developer.work.weixin.qq.com/document/path/94665)

# 微信客服 - 接待人员管理

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execKfServicerCreate`|`reqKfServicerCreate`|`respKfServicerCreate`|+|`POST /cgi-bin/kf/servicer/add`|[添加接待人员](https://developer.work.weixin.qq.com/document/path/94646)
`execKfServicerDelete`|`reqKfServicerDelete`|`respKfServicerDelete`|+|`POST /cgi-bin/kf/servicer/del`|[删除接待人员](https://developer.work.weixin.qq.com/document/path/94647)
`execKfServicerList`|`reqKfServicerList`|`respKfServicerList`|+|`GET /cgi-bin/kf/servicer/list`|[获取接待人员列表](https://developer.work.weixin.qq.com/document/path/94645)

# 微信客服 - 会话分配与消息收发

## API calls

Name|Request Type|Response Type|Access Token|URL|Doc
:---|------------|-------------|------------|:--|:--
`execKfServiceStateGet`|`reqKfServiceStateGet`|`respKfServiceStateGet`|+|`POST /cgi-bin/kf/service_state/get`|[获取会话状态](https://developer.work.weixin.qq.com/document/path/94669)
`execKfServiceStateTrans`|`reqKfServiceStateTrans`|`respKfServiceStateTrans`|+|`POST /cgi-bin/kf/service_state/trans`|[变更会话状态](https://developer.work.weixin.qq.com/document/path/94669)
`execKfSyncMsg`|`reqKfSyncMsg`|`respKfSyncMsg`|+|`POST /cgi-bin/kf/sync_msg`|[读取消息](https://developer.work.weixin.qq.com/document/path/94670)
`execKfSend`|`reqMessage`|`respMessageSend`|+|`POST /cgi-bin/kf/send_msg`|[发送消息](https://developer.work.weixin.qq.com/document/path/94677)
`execKfOnEventSend`|`reqMessage`|`respMessageSend`|+|`POST /cgi-bin/kf/send_msg_on_event`|[发送欢迎语等事件响应消息](https://developer.work.weixin.qq.com/document/path/95122)


# 文档 - 文档管理

## API calls
| Name                          | Request Type                   | Response Type                | Access Token | URL                                            | Doc                                                              |
|:------------------------------|--------------------------------|------------------------------|--------------|:-----------------------------------------------|:-----------------------------------------------------------------|
| `execWedocCreateDoc`          | `reqWedocCreateDoc`                 | `respWedocCreateDoc`         | +            | `POST /cgi-bin/wedoc/create_doc`               | [新建文档](https://developer.work.weixin.qq.com/document/path/97460) |
| `execWedocBatchUpdate`        | `reqWedocBatchUpdate` | `respWedocBatchUpdate`       | +            | `POST /cgi-bin/wedoc/spreadsheet/batch_update` | [新建文档](https://developer.work.weixin.qq.com/document/path/97628) |
| `execWedocGetSheetRangeData`  | `reqWedocGetSheetRangeData` | `respWedocGetSheetRangeData` | +            | `POST /cgi-bin/wedoc/spreadsheet/get_sheet_range_data` | [新建文档](https://developer.work.weixin.qq.com/document/path/97661) |
| `execWedocGetSheetProperties` | `reqWedocGetSheetProperties` | `respWedocGetSheetProperties` | +            | `POST /cgi-bin/wedoc/spreadsheet/get_sheet_properties` | [新建文档](https://developer.work.weixin.qq.com/document/path/97711) |
