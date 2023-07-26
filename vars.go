/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose-zap
 * @Date:        2023-07-26 11:23
 * @Description:
 */

package rzap

const (
	DefaultSingleFilename     = "logs/rzap.log"
	DefaultMultiFilenameInfo  = "logs/info.log"
	DefaultMultiFilenameError = "logs/error.log"
)

type OTP int8

const (
	OutTypeConsole OTP = iota
	OutTypeSingleFileDefault
	OutTypeSingleFileCustom
	OutTypeMultiFileDefault
	OutTypeMultiFileCustom
	OutTypeInfoError
)
