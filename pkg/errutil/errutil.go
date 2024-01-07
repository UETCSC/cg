package errutil

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type error interface {
	Error() string
}

var errorMessages = map[string]string{
	"InvalidAccessKeyId.NotFound": "Current access key are invalid",
	"Message: The specified parameter \"SecurityToken.Expired\" is not valid.": "Current SecurityToken has expired",
	"ErrorCode: InvalidSecurityToken.Expired":                                  "Current SecurityToken has expired",
	"Message: The Access Key is disabled.":                                     "The Access Key is disabled",
	"ErrorCode: Forbidden.RAM":                                                 "Current Access Key do not have permission to execute commands",
	"ErrorCode: NoPermission":                                                  "Current Access Key do not have permission to take over the console)",
	"ErrorCode=NoSuchKey":                                                      "There is no such key in the bucket)",
	"Code=ResourceNotFound, Message=未查询到对应机器":                                  "指定资源不存在 (Resource not found)",
	//Code=UnauthorizedOperation: "Insufficient Access Key permissions are currently available",
	"you are not authorized to perform operation (tat:CreateCommand)": "This Access Key does not have permission to execute commands",
	"network is unreachable":       "Network is unreachable",
	"InvalidSecurityToken.Expired": "STS token has expired",
	"InvalidAccessKeyId.Inactive":  "The current AccessKeyId is inactive",
	"interrupt":                    "程序已退出 (Program exited.)",
	"ErrorCode=AccessDenied, ErrorMessage=\"The bucket you access does not belong to you.\"": "获取 Bucket 信息失败，访问被拒绝 (Failed to get Bucket information, access is denied.)",
	"ExpiredToken":                                                        "当前访问密钥已过期 (Current token has expired)",
	"read: connection reset by peer":                                      "网络连接出现错误，请检查您的网络环境是否正常 (There is an error in your network connection, please check if your network environment is normal.)",
	"Code=ResourceUnavailable.AgentNotInstalled":                          "Agent 未安装 (Agent not installed)",
	"Incorrect IAM authentication information":                            "当前 AK 信息无效 (Current AccessKey information is invalid)",
	"The API does not exist or has not been published in the environment": "当前用户已存在，请指定其他用户名 (User already exists, please specify another user name)",
	"Status=403 Forbidden, Code=AccessDenied":                             "当前权限不足 (Insufficient permissions)",
	"Message: The specified InstanceId does not exist.":                   "指定的实例不存在 (The specified instance does not exist.)",
	"Message: Specified account name already exists in this instance.":    "用户名已存在，请指定其他的用户名 (The username already exists. Please specify a different username.)",
	"Message: Other endpoint exist.":                                      "数据库已经是公开访问的 (The database is already publicly accessible.)",
	"Message: Current DB instance state does not support this operation.": "当前数据库状态不支持此操作，请稍后重试 (The current database state does not support this operation. Please try again later.)",
	"Message: Specified SecurityIPList is not valid.":                     "指定的白名单无效，请检查后重试，注意是否有使用 CIDR 格式 (The specified whitelist is invalid. Please check and try again, ensuring that you are using the CIDR format if required.)",
	"Message: Invalid security ip list specified, duplicated.":            "该白名单列表已存在 (The whitelist entry already exists.)",
}

var errorMessagesNoExit = map[string]string{
	"ErrorCode: Forbidden.RAM": "Current Access Key do not have permission to execute commands",
	//"ErrorCode: Forbidden": "Current Access Key do not have read access to RDS",
	"You are forbidden to list buckets.":                                 "OSS data is not available because the current credential does not have read access to OSS.",
	"ErrorCode: EntityAlreadyExists.User.Policy":                         "Console has been taken over",
	"ErrorCode: EntityAlreadyExists.User":                                "Console has been taken over",
	"ErrorCode: EntityNotExist.User":                                     "Console has been de-taken over",
	"Code=ResourceNotFound, Message=指定资源":                                "ResourceNotFound",
	"InvalidParameter.SubUserNameInUse":                                  "Console has been taken over",
	"you are not authorized to perform operation (cwp:DescribeMachines)": "当前 AK 没有 CWP 权限",
}

var errorMessagesExit = map[string]string{
	"ErrorCode: Forbidden.RAM":     "Current Access Key do not have permission to execute commands",
	"ErrorCode: NoPermission":      "Current Access Key do not have permission to take over the console",
	"network is unreachable":       "Network is unreachable",
	"InvalidSecurityToken.Expired": "STS token has expired",
	"InvalidAccessKeyId.Inactive":  "The current AccessKeyId is inactive",
	//"Message = unauthorized operation, please check the CAM policy. "
	"Code=AuthFailure.SecretIdNotFound": "SecretId does not exist, please enter the correct key.",
	"Code=AuthFailure.SignatureFailure": "Request signature verification failed, please check if your access key is correct.)",
	"read: connection reset by peer":    "There is an error in your network connection, please check if your network environment is normal.)",
	"InvalidAccessKeyId.NotFound":       "Current access key are invalid)",
	"InvalidAccessKeySecret":            "Invalid AccessKey",
}

// HandleErr logs the error and exits. If there are errors in the error string it will check if it contains any of the error keys that are to be logged and if so they will be logged as a warning to the user
// 
// @param e - error to log and
func HandleErr(e error) {
	// This function logs the error messages and exits with an exit code.
	if e != nil {
		log.Traceln(e.Error())
		// Prints a message with a warning message if any of the error messages contain a key.
		for k, v := range errorMessages {
			// If the error is a string of the form e. Error k then exit with an error message.
			if strings.Contains(e.Error(), k) {
				log.Errorln(v)
				os.Exit(0)
			}
		}
		log.Errorln(e)
	}
}

// HandleErrNoExit handles errors that do not exit. This is used to avoid logging in debug mode
// 
// @param e - error to handle or
func HandleErrNoExit(e error) {
	// log all errors and exit
	if e != nil {
		log.Traceln(e.Error())
		// log. Debugln if any of the error messages are not in errorMessagesNoExit
		for k, v := range errorMessagesNoExit {
			// Log a debug message if the error is a string
			if strings.Contains(e.Error(), k) {
				log.Debugln(v)
			}
		}
		// This function will log the error messages and exit with an exit code 0.
		for k, v := range errorMessagesExit {
			// If the error is a string of the form e. Error k then exit with an error message.
			if strings.Contains(e.Error(), k) {
				log.Errorln(v)
				os.Exit(0)
			}
		}
	}
}
